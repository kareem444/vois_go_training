package users

import (
	"net/http"

	"example.com/test/core/mongoDB"
	"example.com/test/core/req"
	"example.com/test/core/res"
	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	body, err := req.Body[User](c, true)

	if err != nil {
		return
	}

	_, exist, checkError := mongoDB.FindOne[any]("user", map[string]any{
		"email": body.Email,
	})

	if checkError != nil {
		res.Error(c, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    checkError.Error(),
		})
		return
	}

	if exist {
		res.Error(c, res.Response{
			StatusCode: http.StatusConflict,
			Message:    "Email already exists",
		})
		return
	}

	id, createError := mongoDB.Insert("user", body)

	if createError != nil {
		res.Error(c, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    createError.Error(),
		})
		return
	}

	body.ID = id

	token, tokenError := _createToken(c, authPayload{
		ID:    body.ID,
		Email: body.Email,
	})

	if tokenError != nil {
		return
	}

	body.Token = token

	res.Success(c, res.Response{
		Message:    "User created successfully",
		StatusCode: http.StatusCreated,
		Data:       body,
	})
}

func login(c *gin.Context) {
	body, err := req.Body[User](c, true)

	if err != nil {
		return
	}

	user, exists, findError := mongoDB.FindOne[User]("user", map[string]any{
		"email": body.Email,
	})

	if findError != nil {
		res.Error(c, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    findError.Error(),
		})
		return
	}

	if !exists {
		res.Error(c, res.Response{
			StatusCode: http.StatusNotFound,
			Message:    "User not found",
		})
		return
	}

	if user.Password != body.Password {
		res.Error(c, res.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid password",
		})
		return
	}

	token, tokenError := _createToken(c, authPayload{
		ID:    user.ID,
		Email: user.Email,
	})

	if tokenError != nil {
		return
	}

	user.Token = token

	res.Success(c, res.Response{
		Message:    "Login successful",
		StatusCode: http.StatusOK,
		Data:       user,
	})
}

func profile(c *gin.Context) {
	auth, _ := req.Auth(c)

	user, exists, err := mongoDB.FindOne[User]("user", map[string]any{
		"email": auth.Email,
	})

	if err != nil {
		res.Error(c, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	if !exists {
		res.Error(c, res.Response{
			StatusCode: http.StatusNotFound,
			Message:    "User not found",
		})
		return
	}

	res.Success(c, res.Response{
		StatusCode: http.StatusOK,
		Data:       user,
	})
}

func findOne(c *gin.Context) {
	id := c.Param("id")

	user, exists, err := mongoDB.FindByID[User]("user", id)

	if err != nil {
		res.Error(c, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	if !exists {
		res.Error(c, res.Response{
			StatusCode: http.StatusNotFound,
			Message:    "User not found",
		})
		return
	}

	res.Success(c, res.Response{
		StatusCode: http.StatusOK,
		Data:       user,
	})
}

func findAll(c *gin.Context) {
	data, err := mongoDB.Find[User]("user")

	if err != nil {
		res.Error(c, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	res.Success(c, res.Response{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}
