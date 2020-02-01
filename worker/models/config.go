package models

import "time"

// Server contains the structure for the server configuration.
type Server struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// Queue contains the configuration for RabbitMQ.
type Queue struct {
	ExchangeName string
	QueueName    string
}
