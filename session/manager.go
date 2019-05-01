package session

import (
	crand "crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	cookieName    = "session_id"
	oneDaySec     = 86400 // 60 * 60 * 24
	sessionMaxAge = oneDaySec * 7
	cookiePath    = "/"
)

func StartNewSession(c *gin.Context, data string) (Session, error) {
	sessionID := generateSessionID()
	if sessionID == "" {
		return Session{}, errors.New("failed to generate session ID")
	}

	storeSessionIDInCookie(c, sessionID, sessionMaxAge)

	session := Session{
		SessionID: sessionID,
		Data:      data,
	}
	return Create(&session)
}

func CurrentSessionID(c *gin.Context) (string, error) {
	return c.Cookie(cookieName)
}

func CurrentSessionData(c *gin.Context) (string, error) {
	sessionID, err := CurrentSessionID(c)
	if err != nil {
		return "", err
	}

	session, err := FindBySessionID(sessionID)
	if err != nil {
		return "", err
	}

	// セッションの改竄をしていなければここは通らないはず
	if isSessionExpired(session) {
		DestroySession(c)
		return "", errors.New("session has expired")
	}

	return GetDataBySessionID(sessionID)
}

func isSessionExpired(session Session) bool {
	expiredAt := session.CreatedAt.Add(sessionMaxAge)
	return time.Now().After(expiredAt)
}

func DestroySession(c *gin.Context) error {
	sessionID, err := CurrentSessionID(c)
	if err != nil {
		return err
	}

	storeSessionIDInCookie(c, sessionID, -1)

	return DeleteBySessionID(sessionID)
}

func storeSessionIDInCookie(c *gin.Context, sessionID string, maxAge int) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     cookieName,
		Value:    url.QueryEscape(sessionID),
		Path:     cookiePath,
		MaxAge:   maxAge,
		Secure:   c.Request.URL.Scheme == "https",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}

func generateSessionID() string {
	b := make([]byte, 64)
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}

	str := base64.URLEncoding.EncodeToString(b) // length: 88
	return str
}
