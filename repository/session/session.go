package session

import (
	"github.com/nekonenene/gin_quiz_app/common"
	"github.com/nekonenene/gin_quiz_app/model"
)

type Session model.Session

func FindBySessionID(sessionID string) (Session, error) {
	db := common.GetDB()
	var session Session
	err := db.Where("session_id = ?", sessionID).First(&session).Error

	return session, err
}

func Create(session *Session) (Session, error) {
	db := common.GetDB()
	err := db.Create(session).Error

	return *session, err
}

func DeleteBySessionID(sessionID string) error {
	db := common.GetDB()
	return db.Where("session_id = ?", sessionID).Delete(Session{}).Error
}

func (sess *Session) Decode() (Data, error) {
	return Decode(sess.EncodedData)
}
