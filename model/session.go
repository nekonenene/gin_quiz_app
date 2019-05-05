package model

import (
	"time"
)

type Session struct {
	SessionID   string    `gorm:"primary_key"                                      json:"session_id"   binding:"required,max=255"`
	EncodedData string    `gorm:"type:text; not null"                              json:"encoded_data" binding:"required,max=65535"`
	CreatedAt   time.Time `gorm:"type:datetime(3); not null; index:created_at_idx" json:"created_at"`
}
