package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nekonenene/quiz_app/common"
	"github.com/nekonenene/quiz_app/user"
)

const (
	portNum = 8013
)

func main() {
	db := common.InitDB()
	defer db.Close()

	router := gin.Default()
	api := router.Group("/api")

	user.UserRouter(api.Group("/user"))

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":" + strconv.Itoa(portNum))
}
