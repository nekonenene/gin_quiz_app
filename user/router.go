package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nekonenene/gin_quiz_app/common"
	"github.com/nekonenene/gin_quiz_app/session"
)

func UserRouter(router *gin.RouterGroup) {
	router.GET("/list", listUser)
	router.GET("/show/:id", showUserByID)
	router.POST("/create", createUser)
}

func listUser(c *gin.Context) {
	sessionID := c.Query("session")
	sess, err := session.FindBySessionID(sessionID)
	if err != nil {
		common.ForbiddenErrorResponse(c)
		return
	}

	data, _ := sess.Decode()
	if !(data.UserID > 0) {
		common.ForbiddenErrorResponse(c)
		return
	}

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
	user.Create()

	common.OkResponse(c, "user", user)
}
