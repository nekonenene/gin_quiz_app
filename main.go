package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nekonenene/gin_quiz_app/common"
	gauth "github.com/nekonenene/gin_quiz_app/oauth/google"
	"github.com/nekonenene/gin_quiz_app/root"
	"github.com/nekonenene/gin_quiz_app/session"
	"github.com/nekonenene/gin_quiz_app/user"
)

func main() {
	db := common.InitDB()
	defer db.Close()

	migrateAll()
	setupRouter()
}

func migrateAll() {
	user.AutoMigrate()
	session.AutoMigrate()
}

func setupRouter() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("assets/html/*")
	router.Static("/img", "assets/img")
	router.Static("/js", "assets/js")

	root.RootRouter(router.Group("/"))

	gauth.InitConf()
	gauth.GoogleOAuthRouter(router.Group("/oauth/google"))

	api := router.Group("/api")
	user.UserRouter(api.Group("/user"))

	port := os.Getenv("SERVER_PORT_NUM")
	router.Run(":" + port)
}
