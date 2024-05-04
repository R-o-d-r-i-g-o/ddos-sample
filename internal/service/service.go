package service

import (
	"log"

	"ddos-sample/internal/api"
	"ddos-sample/pkg/scheduler"
)

func StressWebsite() {
	const SECOND_ENV = 10
	const RATE_LIMIT = 10
	const URL = "https://webhook.site/d4d54ea1-c0f6-4f53-9233-a9ff81793a73"

	sem := scheduler.NewScheduler(RATE_LIMIT, SECOND_ENV)
	sem.Run(func(taskID int) {
		log.Printf("[info] running task %d ...\n", taskID)

		if !api.IsWebsiteStillWorking(URL) {
			log.Printf("[warn] website is down.")
		}
	})
}
