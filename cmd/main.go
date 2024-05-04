package main

import (
	"fmt"
	"sync"

	"ddos-sample/pkg/scheduler"
	"ddos-sample/pkg/shutdown"
)

func main() {
	defer func() {
		fmt.Println("Programa finalizado corretamente.")
	}()

	cleanup := func() {
		fmt.Println("Executando cleanup antes do desligamento...")
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go ExecuteConn(&wg)

	shutdown.GracefulShutdown(cleanup)
	wg.Wait()
}

func ExecuteConn(wg *sync.WaitGroup) {
	defer wg.Done()

	const SECOND_ENV = 5
	const RATE_LIMIT = 1

	sem := scheduler.NewScheduler(RATE_LIMIT, SECOND_ENV)
	sem.Run(func(taskID int) {
		fmt.Println("Veio aqui", taskID)
	})
}
