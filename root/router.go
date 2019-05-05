package root

import (
	"github.com/gin-gonic/gin"
	"github.com/nekonenene/gin_quiz_app/repository/session"
)

func RootRouter(router *gin.RouterGroup) {
	router.GET("/", root)
	router.GET("/signout", signout)
}

func root(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func signout(c *gin.Context) {
	session.DestroySession(c)
	c.Redirect(302, "/")
}
