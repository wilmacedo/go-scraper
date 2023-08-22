package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	Timeout  int
	Key      string
	Endpoint string
}

var env *Environment
var defaultEnv = &Environment{
	Timeout:  25,
	Endpoint: "https://newsapi.org/v2",
}

func Load() error {
	err := godotenv.Load(".env")
	if err != nil {
		return ErrFailedToLoadEnv
	}

	envTimeout := os.Getenv("TIMEOUT")
	var timeout int
	if len(envTimeout) == 0 {
		timeout = defaultEnv.Timeout // Default time
	} else {
		convertedTimeout, err := strconv.Atoi(envTimeout)
		if err != nil {
			return ErrConvertTimeout
		}

		timeout = convertedTimeout
	}

	envKey := os.Getenv("API_KEY")
	if len(envKey) == 0 {
		return ErrApiKeyNotFound
	}

	envEndpoint := os.Getenv("API_ENDPOINT")
	endpoint := envEndpoint
	if len(envEndpoint) == 0 {
		endpoint = defaultEnv.Endpoint
	}

	env = &Environment{
		Timeout:  timeout,
		Key:      envKey,
		Endpoint: endpoint,
	}

	return nil
}

func GetEnv() *Environment {
	return env
}
