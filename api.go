package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	portNum = 8013
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + strconv.Itoa(portNum))
}
