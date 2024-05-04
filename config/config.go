package config

import (
	"os"
	"strconv"

	"log"

	"github.com/joho/godotenv"
)

// this constants holds config names.
const (
	_PATH_URL               = "PATH_URL"
	_CONNECTIONS_PER_SECOND = "CONNECTIONS_PER_SECOND"
	_TIME_IN_SECONDS        = "TIME_IN_SECONDS"
)

// Env holds application pre-setted config.
type Env struct {
	ApiURL        string
	ConnPerSecond int
	TimeInSeconds int
}

// LoadEnv triggers config loading, returning it in a obj.
func LoadEnv() (env Env, err error) {
	if err = godotenv.Load(); err != nil {
		log.Print("[error] Error loading .env file:", err)
		return
	}

	env.ApiURL = os.Getenv(_PATH_URL)
	if env.ApiURL == "" {
		log.Print("[error] PATH_URL is missing.")
		return
	}

	env.ConnPerSecond, err = strconv.Atoi(os.Getenv(_CONNECTIONS_PER_SECOND))
	if err != nil {
		log.Print("[error] CONNECTIONS_PER_SECOND is missing.")
		return
	}

	env.TimeInSeconds, err = strconv.Atoi(os.Getenv(_TIME_IN_SECONDS))
	if err != nil {
		log.Print("[error] TIME_IN_SECONDS is missing.")
		return
	}

	return
}
