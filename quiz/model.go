package quiz

import (
	"time"

	"github.com/nekonenene/gin_quiz_app/common"
	"github.com/nekonenene/gin_quiz_app/user"
)

type Quiz struct {
	ID            uint64    `gorm:"primary_key"                        json:"id"`
	Question      string    `gorm:"not null"                           json:"question"        binding:"required,max=255"`
	Choices       []Choice  `gorm:"primary_key"                        json:"choices"`
	AnswerNumber  int       `gorm:"type:tinyint; not null"             json:"-"               binding:"required,gte=1,lte=100"`
	CreatedUser   user.User `gorm:"foreignkey:CreatedUserID"           json:"-"`
	CreatedUserID uint64    `                                          json:"created_user_id"`
	CreatedAt     time.Time `gorm:"type:datetime(3); not null"         json:"-"`
	UpdatedAt     time.Time `gorm:"type:datetime(3); not null"         json:"-"`
}

type Choice struct {
	Quiz   Quiz   `gorm:"foreignkey:QuizID"                  json:"-"`
	QuizID uint64 `gorm:"primary_key"                        json:"quiz_id"`
	Number int    `gorm:"primary_key; type:tinyint"          json:"number"      binding:"required,gte=1,lte=100"`
	Text   string `gorm:"not null"                           json:"text"        binding:"required,max=255"`
}

func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&Quiz{})
	db.AutoMigrate(&Choice{})
}

func FindAll() ([]Quiz, error) {
	db := common.GetDB()
	var quizzes []Quiz
	err := db.Find(&quizzes).Error

	return quizzes, err
}

func FindByID(id uint64) (Quiz, error) {
	db := common.GetDB()
	var quiz Quiz
	err := db.First(&quiz, id).Error

	return quiz, err
}

func (quiz *Quiz) Create() (Quiz, error) {
	db := common.GetDB()
	err := db.Create(quiz).Error

	return *quiz, err
}
