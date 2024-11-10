package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger(c *gin.Context) {
	// Start timer
	start := time.Now()

	// Process request
	c.Next()

	// Stop timer
	end := time.Now()
	latency := end.Sub(start)

	// Get request data
	requestMethod := c.Request.Method
	requestUri := c.Request.RequestURI
	statusCode := c.Writer.Status()
	clientIP := c.ClientIP()
	userAgent := c.Request.UserAgent()

	log.Printf("| %3d | %13v | %15s | %s | %s | %s |",
		statusCode,
		latency,
		clientIP,
		requestMethod,
		requestUri,
		userAgent,
	)
}
