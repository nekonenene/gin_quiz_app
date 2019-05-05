package model

import (
	"github.com/nekonenene/gin_quiz_app/common"
)

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Session{})
}
