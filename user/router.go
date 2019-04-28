package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nekonenene/gin_quiz_app/common"
)

func UserRouter(router *gin.RouterGroup) {
	router.GET("/list", listUser)
	router.GET("/show/:id", showUserByID)
	router.POST("/create", createUser)
}

func listUser(c *gin.Context) {
	users, err := FindAll()
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}

	common.OkResponse(c, "users", users)
}

func showUserByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		common.ErrorResponse(c, "id must be integer")
		return
	}

	user, err := FindByID(uint(idInt))
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}

	common.OkResponse(c, "user", user)
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		common.BadRequestErrorResponse(c, err.Error())
		return
	}
	Create(&user)

	common.OkResponse(c, "user", user)
}
