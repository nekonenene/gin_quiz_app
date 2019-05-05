package user

import (
	"github.com/nekonenene/gin_quiz_app/model"
	"github.com/nekonenene/gin_quiz_app/registry"
)

type User model.User

func FindAll() ([]User, error) {
	db := registry.DB
	var users []User
	err := db.Find(&users).Error

	return users, err
}

func FindByID(id uint64) (User, error) {
	db := registry.DB
	var user User
	err := db.First(&user, id).Error

	return user, err
}

func FindByOpenID(provider string, providerID string) (User, error) {
	db := registry.DB
	var user User
	err := db.Where("provider = ? AND provider_id >= ?", provider, providerID).First(&user).Error

	return user, err
}

func FindBy(column string, value interface{}) ([]User, error) {
	db := registry.DB
	var users []User
	err := db.Where(column+" = ?", value).Find(&users).Error

	return users, err
}

func (user *User) UpdateOneColumn(column string, value interface{}) (User, error) {
	db := registry.DB
	err := db.Model(&user).Update(column, value).Error

	return *user, err
}

func (user *User) Create() (User, error) {
	db := registry.DB
	err := db.Create(user).Error

	return *user, err
}
