package registry

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/nekonenene/gin_quiz_app/registry/db"
	"github.com/nekonenene/gin_quiz_app/registry/oauth/google"
	"golang.org/x/oauth2"
)

const (
	appHost = "http://localhost:8013" // TODO: appHost は環境変数からの取得にしたい
)

var (
	DB                *gorm.DB
	GoogleOAuthConfig *oauth2.Config
)

func Init() {
	DB = db.Init(db.DBConfig{
		User:         os.Getenv("MYSQL_USER"),
		Password:     os.Getenv("MYSQL_PASSWORD"),
		DatabaseName: os.Getenv("MYSQL_DATABASE"),
	})

	GoogleOAuthConfig = google.InitConf(google.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		BaseURL:      appHost,
	})
}
