package shutdown

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// GracefulShutdown waits for program finishes correclty, then it stops.
func GracefulShutdown(cleanup func()) {
	shutdownCh := make(chan struct{})

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup

	go func() {
		<-signalCh
		close(shutdownCh)
	}()

	go func() {
		<-shutdownCh
		cleanup()
	}()

	wg.Wait()
}
