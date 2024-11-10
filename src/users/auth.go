package users

import (
	"net/http"

	"example.com/test/core/conv"
	"example.com/test/core/jwt"
	"example.com/test/core/res"
	"github.com/gin-gonic/gin"
)

type authPayload struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func _createToken(c *gin.Context, payload authPayload) (string, error) {
	token, err := jwt.Create(conv.ToMap(payload))

	if err != nil {
		res.Json(c, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return "", err
	}

	return token, nil
}
