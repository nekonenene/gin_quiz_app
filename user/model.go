package user

import (
	"time"

	"github.com/nekonenene/gin_quiz_app/common"
)

type User struct {
	ID         uint64    `gorm:"primary_key"                        json:"id"`
	Provider   string    `gorm:"not null; unique_index:openid_idx"  json:"-"           binding:"required,max=255"`
	ProviderID string    `gorm:"not null; unique_index:openid_idx"  json:"-"           binding:"required,max=255"`
	Name       string    `gorm:"not null"                           json:"name"        binding:"required,max=255"`
	Email      string    `gorm:"not null"                           json:"email"       binding:"required,max=255"`
	IsAdmin    bool      `                                          json:"is_admin"`
	CreatedAt  time.Time `gorm:"type:datetime(3); not null"         json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:datetime(3); not null"         json:"-"`
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

func (user *User) Create() (User, error) {
	db := common.GetDB()
	err := db.Create(user).Error

	return *user, err
}
