package main

import (
	"fmt"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	str, _ := common.Encrypt([]byte("Bun Bun Hello, YouTube!"), "password")
	fmt.Println(str)
	str, _ = common.Decrypt(str, "password")
	fmt.Println(string(str))

	db := common.InitDB()
	defer db.Close()

	user.AutoMigrate()
	session.AutoMigrate()

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("gin_quiz_app_session", store))
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("assets/html/*")
	router.Static("/img", "assets/img")

	oauth.InitGoogleOAuth()
	oauth.GoogleRouter(router.Group("/oauth/google"))

	api := router.Group("/api")
	user.UserRouter(api.Group("/user"))

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	router.Run(":" + strconv.Itoa(portNum))
}
