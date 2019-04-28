package user

import (
	"time"

	"github.com/nekonenene/gin_quiz_app/common"
)

type User struct {
	ID        uint64    `gorm:"primary_key"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"type:datetime(3); not null"`
	UpdatedAt time.Time `gorm:"type:datetime(3); not null"`
}

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&User{})
}

func FindByID(id uint) (User, error) {
	db := common.GetDB()
	var user User
	err := db.First(&user, id).Error

	return user, err
}

func Create(user *User) {
	db := common.GetDB()
	db.NewRecord(&user)
	db.Create(user)
}
