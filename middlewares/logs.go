package middlewares

import (
	"log"

	"example.com/test/src/requests_log"
	"github.com/gin-gonic/gin"
)

func Logs(c *gin.Context) {
	c.Next()

	err := requests_log.Create(c)

	if err != nil {
		log.Println("Error creating request log", err)
	}
}
