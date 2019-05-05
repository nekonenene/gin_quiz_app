package main

import (
	"github.com/nekonenene/gin_quiz_app/common"
	"github.com/nekonenene/gin_quiz_app/model"
	gauth "github.com/nekonenene/gin_quiz_app/oauth/google"
	"github.com/nekonenene/gin_quiz_app/router"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	gauth.InitConf()

	model.AutoMigrate()
	router.InitRouter()
}
