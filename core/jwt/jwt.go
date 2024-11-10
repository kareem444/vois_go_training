package jwt

import (
	"errors"
	"time"

	"example.com/test/core/env"
	"github.com/golang-jwt/jwt/v5"
)

func validToken(token string) bool {
	return token != "" && len(token) > 7 && token[:7] == "Bearer "
}

func Create(payload map[string]any) (string, error) {
	JWT_SECRET := env.Get("JWT_SECRET")

	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	for key, value := range payload {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}

	tokenString = "Bearer " + tokenString

	return tokenString, nil
}

func Verify(tokenString string) (map[string]any, error) {
	valid := validToken(tokenString)

	if !valid {
		return nil, errors.New("invalid token")
	}

	tokenString = tokenString[7:]

	JWT_SECRET := env.Get("JWT_SECRET")

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}
