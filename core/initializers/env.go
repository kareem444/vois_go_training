package initializers

import (
	"example.com/test/core/res"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	res.Panic(err)
}
