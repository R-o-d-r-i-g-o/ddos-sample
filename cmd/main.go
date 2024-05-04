package main

import (
	"fmt"
	"sync"
	"time"

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
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("Trabalhando...")
			time.Sleep(time.Second)
		}
	}()

	shutdown.GracefulShutdown(cleanup)
	wg.Wait()
}
