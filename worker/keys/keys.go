package keys

import (
	"os"
	"sync"

	"github.com/jz222/rest-api-with-job-queue/worker/models"

	"github.com/joho/godotenv"
)

var instance *models.Keys
var once sync.Once

func loadEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load environment variables with error:" + err.Error())
	}

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	panic("The environment variable " + key + " was not set.")
}

var envVariables = models.Keys{
	RabbitMQURI: loadEnv("RABBITMQ_URI"),
	Port:        loadEnv("PORT"),
}

// Load initializes all environment variables.
func Load() {
	once.Do(func() {
		instance = &envVariables
	})
}

// Get eturns all environment variables.
func Get() *models.Keys {
	return instance
}
