package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("assets/html/*")
	router.Static("/img", "assets/img")
	router.Static("/js", "assets/js")

	rootRouter(router.Group("/"))

	port := os.Getenv("SERVER_PORT_NUM")
	router.Run(":" + port)
}
