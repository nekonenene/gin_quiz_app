package session

import (
	"time"

	"github.com/nekonenene/gin_quiz_app/model"
	"github.com/nekonenene/gin_quiz_app/registry"
)

type Session model.Session

func FindBySessionID(sessionID string) (Session, error) {
	db := registry.DB
	var session Session
	err := db.Where("session_id = ?", sessionID).First(&session).Error

	return session, err
}

func Create(session *Session) (Session, error) {
	db := registry.DB
	err := db.Create(session).Error

	return *session, err
}

func DeleteBySessionID(sessionID string) error {
	db := registry.DB
	return db.Where("session_id = ?", sessionID).Delete(Session{}).Error
}

func (sess *Session) Decode() (Data, error) {
	return Decode(sess.EncodedData)
}

// 期限切れのセッションか判定
func (sess *Session) IsExpired(maxAgeSec int) bool {
	expiredAt := sess.CreatedAt.Add(time.Second * time.Duration(maxAgeSec))
	return time.Now().After(expiredAt)
}
