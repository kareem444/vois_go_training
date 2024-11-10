package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
	Success    bool        `json:"success"`
}

func Json(c *gin.Context, response Response) {
	handleRepose := gin.H{
		"success": response.Success,
	}

	if response.Message != "" {
		handleRepose["message"] = response.Message
	}

	if response.Data != nil {
		handleRepose["data"] = response.Data
	}

	if response.StatusCode != 0 {
		handleRepose["statusCode"] = response.StatusCode
	}

	c.JSON(response.StatusCode, handleRepose)
}

func Success(c *gin.Context, response Response) {
	if response.StatusCode == 0 {
		response.StatusCode = http.StatusOK
	}

	if response.Message == "" {
		response.Message = "Success"
	}

	response.Success = true

	Json(c, response)
}

func Error(c *gin.Context, response Response) {
	if response.StatusCode == 0 {
		response.StatusCode = http.StatusBadRequest
	}

	if response.Message == "" {
		response.Message = "Bad Request"
	}

	response.Success = false

	Json(c, response)
}
