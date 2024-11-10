package articles

import (
	"example.com/test/middlewares"
	"github.com/gin-gonic/gin"
)

func Controller(r *gin.RouterGroup) {
	route := r.Group("articles")

	route.POST("/inspect", inspect)

	route.POST("/", middlewares.Authentication, create)
	route.GET("/", middlewares.Authentication, findAll)
	route.GET("/paginate", middlewares.Authentication, paginate)
}
