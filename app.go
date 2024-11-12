package main

import (
	"example.com/test/schedules"
	"example.com/test/src/articles"
	"example.com/test/src/requests_log"
	"example.com/test/src/users"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func registerControllers(routes *gin.RouterGroup) {
	users.Controller(routes)
	articles.Controller(routes)
	requests_log.Controller(routes)
}

func registerCron() {
	var c *cron.Cron = cron.New()
	schedules.UpdateArticles(c)
	schedules.SaveLogs(c)

	c.Start()
}
