package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OkResponse(c *gin.Context, key string, obj interface{}) {
	c.JSON(200, gin.H{
		key: obj,
	})
}

// 500 error
func ErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": message,
	})
}

// 400 error
func BadRequestErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": message,
	})
}

// 403 error
func ForbiddenErrorResponse(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"error": "Forbidden error",
	})
}
