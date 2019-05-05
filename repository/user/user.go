package user

import (
	"github.com/nekonenene/gin_quiz_app/common"
	"github.com/nekonenene/gin_quiz_app/model"
)

type User model.User

func FindAll() ([]User, error) {
	db := common.GetDB()
	var users []User
	err := db.Find(&users).Error

	return users, err
}

func FindByID(id uint64) (User, error) {
	db := common.GetDB()
	var user User
	err := db.First(&user, id).Error

	return user, err
}

func FindByOpenID(provider string, providerID string) (User, error) {
	db := common.GetDB()
	var user User
	err := db.Where("provider = ? AND provider_id >= ?", provider, providerID).First(&user).Error

	return user, err
}

func FindBy(column string, value interface{}) ([]User, error) {
	db := common.GetDB()
	var users []User
	err := db.Where(column+" = ?", value).Find(&users).Error

	return users, err
}

func (user *User) UpdateOneColumn(column string, value interface{}) (User, error) {
	db := common.GetDB()
	err := db.Model(&user).Update(column, value).Error

	return *user, err
}

func (user *User) Create() (User, error) {
	db := common.GetDB()
	err := db.Create(user).Error

	return *user, err
}
