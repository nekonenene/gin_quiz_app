package main

import (
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/nekonenene/gin_quiz_app/common"
	"github.com/nekonenene/gin_quiz_app/oauth"
	"github.com/nekonenene/gin_quiz_app/user"
)

const (
	portNum = 8013
)

func main() {
	db := common.InitDB()
	defer db.Close()

	user.AutoMigrate()

	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("gin_quiz_app_session", store))
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("assets/html/*")

	oauth.InitGoogleOAuth()
	oauth.GoogleRouter(router.Group("/oauth/google"))

	api := router.Group("/api")
	user.UserRouter(api.Group("/user"))

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	router.Run(":" + strconv.Itoa(portNum))
}
