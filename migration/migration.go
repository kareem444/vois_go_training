package main

import (
	"example.com/test/core/initializers"
	"example.com/test/core/logger"
	"example.com/test/src/articles"
)

func init() {
	logger.Info("Migrate Initializing")
	initializers.LoadEnv()
	initializers.InitDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&articles.Article{})
	logger.FatalWithMessage("Error migrating: ", err)
	logger.Info("Migrate Finished Successfully")
}
