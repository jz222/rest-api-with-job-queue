package config

import (
	"time"

	"github.com/jz222/rest-api-with-job-queue/worker/models"
)

// Server contains the server configuration.
var Server = models.Server{
	ReadTimeout:  10 * time.Second,
	WriteTimeout: 10 * time.Second,
	IdleTimeout:  10 * time.Second,
}

// Queue contains the RabbitMQ configuration.
var Queue = models.Queue{
	ExchangeName: "events",
	QueueName:    "test",
}
