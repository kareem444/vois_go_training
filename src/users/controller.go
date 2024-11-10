package users

import (
	"example.com/test/middlewares"
	"github.com/gin-gonic/gin"
)

func Controller(r *gin.RouterGroup) {
	route := r.Group("users")

	route.GET("/", middlewares.Authentication, findAll)
	route.GET("/:id", middlewares.Authentication, findOne)

	route.POST("/register", register)
	route.POST("/login", login)
	route.GET("/profile", middlewares.Authentication, profile)
}
