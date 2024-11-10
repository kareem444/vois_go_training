package req

import (
	"example.com/test/core/conv"
	"github.com/gin-gonic/gin"
)

type auth struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func Auth(c *gin.Context) (data auth, exist bool) {
	request, exist := Req[map[string]any](c, "auth")

	if !exist {
		return data, false
	}

	authData := request
	data.ID = conv.ToString(authData["id"])
	data.Email = conv.ToString(authData["email"])

	return data, true
}
