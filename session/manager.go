package session

import (
	crand "crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

const (
	cookieName   = "session_id"
	oneDaySec    = 86400 // 60 * 60 * 24
	cookieMaxAge = oneDaySec * 7
	cookiePath   = "/"
)

func StartNewSession(c *gin.Context, data string) (Session, error) {
	sessionID := generateSessionID()
	if sessionID == "" {
		return Session{}, errors.New("failed to generate session ID")
	}

	storeSessionIDInCookie(c, sessionID)

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

	return GetDataBySessionID(sessionID)
}

func storeSessionIDInCookie(c *gin.Context, sessionID string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     cookieName,
		Value:    url.QueryEscape(sessionID),
		MaxAge:   cookieMaxAge,
		Secure:   c.Request.URL.Scheme == "https",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
}

func generateSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}

	str := base64.URLEncoding.EncodeToString(b) // length: 44
	return str
}
