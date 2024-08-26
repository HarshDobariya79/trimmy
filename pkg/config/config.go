package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Env is the exported map that stores the environment variables.
var Env map[string]string

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading environment variables.")
	}

	Env = make(map[string]string)

	envVariables := []string{
		"REDIS_HOST",
		"REDIS_PASSWORD",
		"BACKEND_HOST",
	}

	for _, key := range envVariables {
		value := os.Getenv(key)
		if value == "" {
			log.Printf("Warning: %s environment variable is not set.", key)
		}
		Env[key] = value
	}
}
