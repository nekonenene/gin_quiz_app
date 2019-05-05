package model

import (
	"time"
)

type User struct {
	ID         uint64    `gorm:"primary_key"                        json:"id"`
	Provider   string    `gorm:"not null; unique_index:openid_idx"  json:"-"`
	ProviderID string    `gorm:"not null; unique_index:openid_idx"  json:"-"`
	Name       string    `gorm:"not null"                           json:"name"        binding:"required,max=255"`
	Email      string    `gorm:"not null"                           json:"email"       binding:"required,max=255,email"`
	IsAdmin    bool      `                                          json:"-"`
	CreatedAt  time.Time `gorm:"type:datetime(3); not null"         json:"-"`
	UpdatedAt  time.Time `gorm:"type:datetime(3); not null"         json:"-"`
}
