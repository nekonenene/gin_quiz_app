package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	router.GET("/list", listUser)
	router.GET("/show/:id", showUserByID)
	router.POST("/create", createUser)
}

func listUser(c *gin.Context) {
	users, err := FindAll()
	if err != nil {
		errorResponse(c, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func showUserByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		errorResponse(c, "id must be integer")
		return
	}

	user, err := FindByID(uint(idInt))
	if err != nil {
		errorResponse(c, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func createUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	Create(&user)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func errorResponse(c *gin.Context, message string) {
	c.JSON(500, gin.H{
		"error": message,
	})
}
