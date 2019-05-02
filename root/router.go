package root

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nekonenene/gin_quiz_app/session"
)

func RootRouter(router *gin.RouterGroup) {
	router.GET("/", root)
	router.GET("/home", home)
	router.GET("/signout", signout)
}

func root(c *gin.Context) {
	if session.IsSignin(c) {
		c.Redirect(302, "/home")
	}

	c.HTML(200, "index.html", gin.H{})
}

func home(c *gin.Context) {
	userID, err := session.CurrentUserID(c)
	if err != nil {
		c.Redirect(302, "/")
	}

	c.HTML(200, "home.html", gin.H{
		"userID": strconv.FormatUint(userID, 10),
	})
}

func signout(c *gin.Context) {
	session.DestroySession(c)
	c.Redirect(302, "/")
}
