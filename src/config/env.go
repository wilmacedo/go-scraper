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

func Load() error {
	err := godotenv.Load(".env")
	if err != nil {
		return ErrFailedToLoadEnv
	}

	envTimeout := os.Getenv("TIMEOUT")
	timeout, err := strconv.Atoi(envTimeout)
	if err != nil {
		return ErrConvertTimeout
	}

	env = &Environment{
		Timeout: timeout,
	}

	return nil
}

func GetEnv() *Environment {
	return env
}
