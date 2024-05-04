package main

import (
	"fmt"
	"sync"

	"ddos-sample/internal/service"
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
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		service.StressWebsite()
	}(&wg)

	shutdown.GracefulShutdown(cleanup)
	wg.Wait()
}
