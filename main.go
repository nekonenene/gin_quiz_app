package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nekonenene/gin_quiz_app/common"
	gauth "github.com/nekonenene/gin_quiz_app/oauth/google"
	"github.com/nekonenene/gin_quiz_app/session"
	"github.com/nekonenene/gin_quiz_app/user"
)

const (
	portNum = 8013
)

func main() {
	// TODO: remove this debug code
	str, _ := common.Encrypt([]byte("Bun Bun Hello, YouTube!"), "password")
	log.Println(str)
	bytes, _ := common.Decrypt(str, "password")
	log.Println(string(bytes))

	db := common.InitDB()
	defer db.Close()

	user.AutoMigrate()
	session.AutoMigrate()

	router := gin.Default()
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("assets/html/*")
	router.Static("/img", "assets/img")

	gauth.InitConf()
	gauth.GoogleOAuthRouter(router.Group("/oauth/google"))

	api := router.Group("/api")
	user.UserRouter(api.Group("/user"))

	router.GET("/", func(c *gin.Context) {
		// TODO: remove this debug code
		userID, _ := session.CurrentUserID(c)
		log.Printf("user ID: %v\n", userID)

		c.HTML(200, "index.html", gin.H{})
	})

	router.Run(":" + strconv.Itoa(portNum))
}
