package service

import (
	"log"

	"ddos-sample/config"
	"ddos-sample/internal/api"
	"ddos-sample/pkg/scheduler"
)

// Service holds public api methods.
type Service interface {
	StressWebsite()
}

// service holds its conn info.
type service struct {
	env config.Env
	api api.Api
}

// NewService starts service layer.
func NewService(env config.Env, api api.Api) Service {
	return &service{env, api}
}

// StressWebsite call website based on configured info.
func (s *service) StressWebsite() {
	scheduler := scheduler.NewScheduler(s.env.ConnPerSecond, s.env.TimeInSeconds)

	scheduler.Run(func(taskID int) {
		log.Printf("[info] running task %d ...\n", taskID)

		if !s.api.IsWebsiteStillWorking(s.env.ApiURL) {
			log.Printf("[warn] website is down.")
		}
	})
}
