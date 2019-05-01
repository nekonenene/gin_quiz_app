package session

import (
	crand "crypto/rand"
	"encoding/base64"
	"io"
	"time"

	"github.com/nekonenene/gin_quiz_app/common"
)

type Session struct {
	SessionID string    `gorm:"primary_key"                                      json:"session_id"  binding:"required,max=255"`
	Data      string    `gorm:"type:text; not null"                              json:"name"        binding:"required,max=255"`
	CreatedAt time.Time `gorm:"type:datetime(3); not null; index:created_at_idx" json:"created_at"`
}

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&Session{})
}

func FindByID(session_id string) (Session, error) {
	db := common.GetDB()
	var session Session
	err := db.Where("session_id = ?", session_id).First(&session).Error

	return session, err
}

func Create(session *Session) (Session, error) {
	db := common.GetDB()
	err := db.Create(session).Error

	return *session, err
}

func CreateWithData(data string) (Session, error) {
	sessionID, err := generateSessionId()
	if err != nil {
		return Session{}, err
	}

	session := Session{
		SessionID: sessionID,
		Data:      data,
	}
	return Create(&session)
}

func generateSessionId() (string, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return "", err
	}

	str := base64.URLEncoding.EncodeToString(b) // length: 44
	return str, nil
}
