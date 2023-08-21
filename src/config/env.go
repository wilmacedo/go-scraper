package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	Timeout int
}

var env *Environment
var defaultEnv = &Environment{
	Timeout: 25,
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

	env = &Environment{
		Timeout: timeout,
	}

	return nil
}

func GetEnv() *Environment {
	return env
}
