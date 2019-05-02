package session

import (
	"time"

	"github.com/nekonenene/gin_quiz_app/common"
)

type Session struct {
	SessionID   string    `gorm:"primary_key"                                      json:"session_id"   binding:"required,max=255"`
	EncodedData string    `gorm:"type:text; not null"                              json:"encoded_data" binding:"required,max=65535"`
	CreatedAt   time.Time `gorm:"type:datetime(3); not null; index:created_at_idx" json:"created_at"`
}

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&Session{})
}

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
