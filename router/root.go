package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nekonenene/gin_quiz_app/repository/session"
	"github.com/nekonenene/gin_quiz_app/router/oauth/google"
	"github.com/nekonenene/gin_quiz_app/router/user"
)

func rootRouter(router *gin.RouterGroup) {
	router.GET("/", root)
	router.GET("/signout", signout)

	google.GoogleOAuthRouter(router.Group("/oauth/google"))

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
