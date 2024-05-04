package scheduler

import (
	"sync"
	"time"
)

// Scheduler holds Runner methods.
type Scheduler interface {
	Run(callback func(taskID int))
}

// scheduler holds a scheduler base fuction info.
type scheduler struct {
	timeInSeconds int
	rateLimit     int
}

// NewScheduler retrives a new SCHEDULER instance.
func NewScheduler(rateLimit, timeInSeconds int) Scheduler {
	return &scheduler{
		timeInSeconds: timeInSeconds,
		rateLimit:     rateLimit,
	}
}

// Run executes callback function in each registered second.
func (s *scheduler) Run(callback func(taskID int)) {
	var wg sync.WaitGroup
	var count int

	closeWait := func(wg *sync.WaitGroup, taskID int) {
		defer wg.Done()
		callback(taskID)
	}

	for count <= s.timeInSeconds {
		for i := 0; i < s.rateLimit; i++ {
			wg.Add(1)
			go closeWait(&wg, count)
		}
		wg.Wait()

		time.Sleep(1 * time.Second)
		count++
	}
}
