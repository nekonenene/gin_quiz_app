package root

import (
	"github.com/gin-gonic/gin"
	"github.com/nekonenene/gin_quiz_app/session"
)

func RootRouter(router *gin.RouterGroup) {
	router.GET("/", root)
	router.GET("/home", home)
	router.GET("/signout", signout)
}

func root(c *gin.Context) {
	userID, err := session.CurrentUserID(c)
	if err == nil && userID > 0 {
		c.Redirect(302, "/home")
	}

	c.HTML(200, "index.html", gin.H{})
}

func home(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{})
}

func signout(c *gin.Context) {
	session.DestroySession(c)
	c.Redirect(302, "/")
}
