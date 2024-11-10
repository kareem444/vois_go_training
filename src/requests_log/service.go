package requests_log

import (
	"net/http"
	"time"

	"example.com/test/core/mongoDB"
	"example.com/test/core/res"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) error {
	requestLog := RequestLog{
		Method:     c.Request.Method,
		URL:        c.Request.RequestURI,
		StatusCode: c.Writer.Status(),
		ClientIP:   c.ClientIP(),
		UserAgent:  c.Request.UserAgent(),
		Timestamp:  time.Now(),
	}

	_, err := mongoDB.Insert("request_logs", requestLog)

	return err
}

func GetLogs() ([]RequestLog, error) {
	data, err := mongoDB.Find[RequestLog]("request_logs")

	if err != nil {
		return nil, err
	}

	return data, nil
}

func findAll(c *gin.Context) {
	data, err := mongoDB.Find[RequestLog]("request_logs")

	if err != nil {
		res.Error(c, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get data",
		})
		return
	}

	res.Success(c, res.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}
