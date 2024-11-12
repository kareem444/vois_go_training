package schedules

import (
	"example.com/test/src/articles"
	"github.com/robfig/cron/v3"
)

func UpdateArticles(x *cron.Cron) {
	x.AddFunc("@every 30m", func() {
		articles.UpdateScheduler()
	})
}
