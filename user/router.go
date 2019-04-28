package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserRouter : router
func UserRouter(router *gin.RouterGroup) {
	router.GET("/list", listUser)
	router.GET("/show/:id", showUserByID)
	router.POST("/create", createUser)
}

func listUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User List!",
	})
}

func showUserByID(c *gin.Context) {
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
}

func createUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	c.JSON(200, gin.H{
		"message": "Created user: " + user.Name,
	})
}
