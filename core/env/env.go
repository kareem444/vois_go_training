package env

import (
	"os"
	"strconv"
)

func Get(key string, defaultValue ...string) string {
	value := os.Getenv(key)
	if value == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return value
}

func GetInt(key string, defaultValue ...int) int {
	value := Get(key)

	intValue, err := strconv.Atoi(value)
	if err != nil && len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return intValue
}
