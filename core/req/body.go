package req

import (
	"net/http"

	"example.com/test/core/res"
	"github.com/gin-gonic/gin"
)

func Body[T any](c *gin.Context, resError ...bool) (T, error) {
	var body T

	if err := c.ShouldBindJSON(&body); err != nil {
		if len(resError) > 0 && resError[0] {
			res.Error(c, res.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			})
		}
		return body, err
	}

	return body, nil
}

func BodyToMap(c *gin.Context, resError ...bool) (map[string]any, error) {
	var body map[string]any

	if err := c.BindJSON(&body); err != nil {
		if len(resError) > 0 && resError[0] {
			res.Error(c, res.Response{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			})
		}
		return body, err
	}

	return body, nil
}
