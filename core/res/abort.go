package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Abort struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func AbortJson(c *gin.Context, abort Abort) {
	method := c.Request.Method
	path := c.Request.URL.Path

	if abort.Message == "" {
		abort.Message = "Internal Server Error"
	}

	if abort.StatusCode == 0 {
		abort.StatusCode = http.StatusInternalServerError
	}

	handleRepose := gin.H{
		"message":    abort.Message,
		"statusCode": abort.StatusCode,
		"method":     method,
		"path":       path,
	}

	c.JSON(abort.StatusCode, handleRepose)
}
