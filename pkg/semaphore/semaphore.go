package semaphore

// Semaphore
type Semaphore interface {
	Acquire()
	Release()
}

// semaphore
type semaphore struct {
	permits chan struct{}
}

// NewSemaphore
func NewSemaphore(permits int) Semaphore {
	return &semaphore{permits: make(chan struct{}, permits)}
}

// Acquire
func (s *semaphore) Acquire() {
	s.permits <- struct{}{}
}

// Release
func (s *semaphore) Release() {
	<-s.permits
}
