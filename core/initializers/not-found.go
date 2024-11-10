package initializers

import (
	"net/http"

	"example.com/test/core/res"
	"github.com/gin-gonic/gin"
)

func RouteNotFound(c *gin.Context) {
	res.AbortJson(c, res.Abort{
		Message:    "Not found Route",
		StatusCode: http.StatusNotFound,
	})
}

func MethodNotAllowed(c *gin.Context) {
	res.AbortJson(c, res.Abort{
		Message:    "Method not allowed",
		StatusCode: http.StatusMethodNotAllowed,
	})
}
