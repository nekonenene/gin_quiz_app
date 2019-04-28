package user

import (
	"time"

	"github.com/nekonenene/gin_quiz_app/common"
)

type User struct {
	ID        uint64    `gorm:"primary_key"                json:"id"`
	Name      string    `gorm:"not null"                   json:"name"`
	CreatedAt time.Time `gorm:"type:datetime(3); not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime(3); not null" json:"-"`
}

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&User{})
}

func FindAll() ([]User, error) {
	db := common.GetDB()
	var users []User
	err := db.Find(&users).Error

	return users, err
}

func FindByID(id uint) (User, error) {
	db := common.GetDB()
	var user User
	err := db.First(&user, id).Error

	return user, err
}

func Create(user *User) (User, error) {
	db := common.GetDB()
	err := db.Create(user).Error

	return *user, err
}
