package middlewares

import (
	"net/http"

	"example.com/test/core/conv"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

var (
	rate, _     = limiter.NewRateFromFormatted("20-M") // 20 requests per minute
	store       = memory.NewStore()
	rateLimiter = limiter.New(store, rate)
)

func Limiter(c *gin.Context) {
	key := c.ClientIP()

	// Get rate limit context for the IP key
	context, err := rateLimiter.Get(c, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Set rate limit headers for client awareness
	c.Header("X-RateLimit-Limit", conv.ToString(context.Limit))
	c.Header("X-RateLimit-Remaining", conv.ToString(context.Remaining))
	c.Header("X-RateLimit-Reset", conv.ToString(context.Reset))
	c.Header("X-Clients-IP", key)

	// Check if the limit has been reached
	if context.Reached {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
		c.Abort()
		return
	}

	c.Next()
}
