package req

import "github.com/gin-gonic/gin"

func Req[T any](c *gin.Context, search string) (data T, exist bool) {
	request, exist := c.Get(search)

	if !exist {
		return data, false
	}

	data = request.(T)
	return data, true
}
