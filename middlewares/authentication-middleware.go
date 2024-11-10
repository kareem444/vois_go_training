package middlewares

import (
	"net/http"

	"example.com/test/core/jwt"
	"example.com/test/core/res"
	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) {
	token := c.GetHeader("Authorization")

	payload, err := jwt.Verify(token)

	if err != nil {
		res.Error(c, res.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized",
		})
		c.Abort()
		return
	}

	c.Set("auth", payload)
	c.Next()
}
