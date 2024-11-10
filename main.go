package main

import (
	"example.com/test/core/initializers"
	"example.com/test/core/logger"
	"example.com/test/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	logger.Info("=== Main Initializing ===")
	initializers.LoadEnv()
	initializers.InitDB()
	initializers.InitMongo()
}

func main() {
	registerCron() // comment while developing

	server := gin.Default()
	server.HandleMethodNotAllowed = true

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	server.Use(middlewares.Logger)
	server.Use(middlewares.Logs) // comment while developing
	server.Use(middlewares.Limiter) // comment while developing

	server.NoRoute(initializers.RouteNotFound)
	server.NoMethod(initializers.MethodNotAllowed)

	routes := server.Group("api/v1/")
	registerControllers(routes)

	server.Run(":8080")
}