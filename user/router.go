package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRouter : router
func UserRouter(router *gin.RouterGroup) {
	router.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "User List!",
		})
	})

	router.POST("/show/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "User ID: " + id,
		})
	})
}
