package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nekonenene/gin_quiz_app/common"
	"github.com/nekonenene/gin_quiz_app/oauth"
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
	str, _ = common.Decrypt(str, "password")
	log.Println(string(str))

	db := common.InitDB()
	defer db.Close()

	user.AutoMigrate()
	session.AutoMigrate()

	router := gin.Default()
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("assets/html/*")
	router.Static("/img", "assets/img")

	oauth.InitGoogleOAuth()
	oauth.GoogleRouter(router.Group("/oauth/google"))

	api := router.Group("/api")
	user.UserRouter(api.Group("/user"))

	router.GET("/", func(c *gin.Context) {
		// TODO: remove this debug code
		data, err := session.CurrentSessionData(c)
		log.Printf("session data: %v, err: %v\n", data, err)
		if data == "" {
			session.StartNewSession(c, "hogehoge")
		}

		c.HTML(200, "index.html", gin.H{})
	})

	router.Run(":" + strconv.Itoa(portNum))
}
