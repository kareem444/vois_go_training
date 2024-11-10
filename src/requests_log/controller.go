package requests_log

import (
	"github.com/gin-gonic/gin"
)

func Controller(r *gin.RouterGroup) {
	route := r.Group("requests_log")

	route.GET("/", findAll)
}
