package session

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nekonenene/gin_quiz_app/common"
)

const (
	cookieName      = "session_id"
	oneDaySec       = 86400 // 60 * 60 * 24
	sessionMaxAge   = oneDaySec * 7
	cookiePath      = "/"
	sessionIDLength = 64
)

// 新しい session ID を作成し、cookie および DB に保存。既存の session は削除する
func StartNewSession(c *gin.Context, data *Data) (Session, error) {
	sessionID := common.RandomString(sessionIDLength)
	if sessionID == "" {
		return Session{}, errors.New("failed to generate session ID")
	}
	encodedData, err := data.Encode()
	if err != nil {
		return Session{}, err
	}

	DestroySession(c)
	storeSessionIDInCookie(c, sessionID)

	session := Session{
		SessionID:   sessionID,
		EncodedData: encodedData,
	}
	return Create(&session)
}

// 現在の session ID を cookie から呼び出す
func CurrentSessionID(c *gin.Context) (string, error) {
	return c.Cookie(cookieName)
}

// 現在のセッションデータを呼び出す
// DB から呼び出した値が不正な場合は session の削除処理をおこなう
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
	if session.IsExpired(sessionMaxAge) {
		DestroySession(c)
		return "", errors.New("session has expired")
	}

	return session.EncodedData, nil
}

// session からユーザーIDを取得する
func CurrentUserID(c *gin.Context) (uint64, error) {
	encoded, err := CurrentSessionData(c)
	if err != nil {
		return 0, err
	}

	data, err := Decode(encoded)
	if err != nil {
		return 0, err
	}

	return data.UserID, nil
}

// ログイン中のユーザーか
func IsSignin(c *gin.Context) bool {
	userID, err := CurrentUserID(c)
	return err == nil && userID > 0
}

// session ID を cookie および DB から削除
func DestroySession(c *gin.Context) error {
	sessionID, err := CurrentSessionID(c)
	if err != nil {
		return err
	}

	deleteSessionIDInCookie(c)

	return DeleteBySessionID(sessionID)
}

// cookie に session ID を保存
func storeSessionIDInCookie(c *gin.Context, sessionID string) {
	common.SetCookie(c, cookieName, sessionID, sessionMaxAge)
}

// cookie の session ID を削除
func deleteSessionIDInCookie(c *gin.Context) {
	common.SetCookie(c, cookieName, "", -1)
}
