package router

import (
	"os"

	"github.com/gin-gonic/gin"
	gauth "github.com/nekonenene/gin_quiz_app/oauth/google"
	"github.com/nekonenene/gin_quiz_app/repository/session"
	"github.com/nekonenene/gin_quiz_app/router/user"
)

func InitRouter() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("assets/html/*")
	router.Static("/img", "assets/img")
	router.Static("/js", "assets/js")

	rootRouter(router.Group("/"))

	port := os.Getenv("SERVER_PORT_NUM")
	router.Run(":" + port)
}

func rootRouter(router *gin.RouterGroup) {
	router.GET("/", root)
	router.GET("/signout", signout)

	gauth.InitConf()
	gauth.GoogleOAuthRouter(router.Group("/oauth/google"))

	api := router.Group("/api")
	user.UserRouter(api.Group("/user"))
}

func root(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func signout(c *gin.Context) {
	session.DestroySession(c)
	c.Redirect(302, "/")
}
