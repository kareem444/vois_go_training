package schedules

import (
	"example.com/test/core/firebase_helper"
	"example.com/test/src/requests_log"
	"github.com/robfig/cron/v3"
)

func SaveLogs(x *cron.Cron) {
	x.AddFunc("@weekly", func() {
		logs, err := requests_log.GetLogs()

		if err != nil {
			return
		}

		firebase_helper.UploadAsJson(logs)
	})

	x.Start()
}
