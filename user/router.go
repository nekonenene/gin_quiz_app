package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserRouter : router
func UserRouter(router *gin.RouterGroup) {
	router.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "User List!",
		})
	})

	router.GET("/show/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, _ := strconv.Atoi(id)
		user, err := FindByID(uint(idInt))
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "User name: " + user.Name,
		})
	})

	router.POST("/create", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)

		c.JSON(200, gin.H{
			"message": "Created user: " + user.Name,
		})
	})
}
