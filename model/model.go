package model

import "github.com/nekonenene/gin_quiz_app/registry"

func AutoMigrate() {
	db := registry.DB

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Session{})
}
