package main

import (
	"log"
	"sync"

	"ddos-sample/config"
	"ddos-sample/internal/api"
	"ddos-sample/internal/service"
	"ddos-sample/pkg/shutdown"
)

// TODO: implement a log-manager after all.

var dependencies deps

// init is triggered before app start.
func init() {
	dependencies = loadDependencies()
}

// main starts app after cli-command operations.
func main() {

	defer func() {
		log.Println("[info] Programa finalizado corretamente.")
	}()

	cleanup := func() {
		log.Println("[info] Executando cleanup antes do desligamento...")
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		dependencies.service.StressWebsite()
	}(&wg)

	shutdown.GracefulShutdown(cleanup)
	wg.Wait()
}

// deps holds depency inversion configs.
type deps struct {
	env     config.Env
	service service.Service
}

// NOTE: function bellow returns summarized info only.
// loadDependencies runns dependecy inversion.
func loadDependencies() (d deps) {
	var err error

	d.env, err = config.LoadEnv()
	if err != nil {
		log.Fatal("[error] env file didn't load correctly, info: ", err.Error())
		return
	}

	d.service = service.NewService(d.env, api.NewAPI())
	return
}
