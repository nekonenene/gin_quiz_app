package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nekonenene/gin_quiz_app/common"
	"github.com/nekonenene/gin_quiz_app/session"
)

func UserRouter(router *gin.RouterGroup) {
	router.GET("/current", currentUser)
	router.POST("/update", updateUser)
	router.GET("/list", listUser)
	router.GET("/show/:id", showUserByID)
	router.POST("/create", createUser)
}

func currentUser(c *gin.Context) {
	uid, err := session.CurrentUserID(c)
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}

	user, err := FindByID(uid)
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}

	common.OkResponse(c, "user", user)
}

func updateUser(c *gin.Context) {
	var requestUser User
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		common.BadRequestErrorResponse(c, err.Error())
		return
	}

	uid, err := session.CurrentUserID(c)
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}

	if uid != requestUser.ID {
		common.BadRequestErrorResponse(c, err.Error())
		return
	}

	user, err := FindByID(uid)
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}

	user, err = user.UpdateOneColumn("name", requestUser.Name)
	if err != nil {
		common.ErrorResponse(c, err.Error())
		return
	}

	common.OkResponse(c, "user", user)
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
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		common.ErrorResponse(c, "id must be integer")
		return
	}

	user, err := FindByID(uid)
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
