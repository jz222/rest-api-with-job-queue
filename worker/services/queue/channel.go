package queue

import (
	"log"
	"time"

	"github.com/jz222/rest-api-with-job-queue/worker/config"
	"github.com/jz222/rest-api-with-job-queue/worker/keys"
	"github.com/streadway/amqp"
)

var chann *amqp.Channel
var conn *amqp.Connection

func cleanup() {
	r := recover()
	if r != nil && conn != nil {
		log.Println("Closing RabbitMQ Connection due to error", r)
		conn.Close()
	}
}

func connect() bool {
	defer cleanup()

	log.Println("Attempting to create a connection to RabbitMQ")

	// Connect to RabbitMQ
	connection, err := amqp.Dial(keys.Get().RabbitMQURI)
	if err != nil {
		panic("Could not establish connection to RabbitMQ:" + err.Error())
	}

	conn = connection

	// Create a channel from the connection
	channel, err := connection.Channel()
	if err != nil {
		panic("Could not open RabbitMQ channel:" + err.Error())
	}

	// Create a new exchange that acts as broker for the queue
	err = channel.ExchangeDeclare(config.Queue.ExchangeName, "topic", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	// Create a new queue
	_, err = channel.QueueDeclare(config.Queue.QueueName, true, false, false, false, nil)
	if err != nil {
		panic("Could not declare queue:" + err.Error())
	}

	// Bind the queue to the exchange
	err = channel.QueueBind(config.Queue.QueueName, "#", config.Queue.ExchangeName, false, nil)
	if err != nil {
		panic("Could not bind queue:" + err.Error())
	}

	chann = channel

	log.Println("Successfully established connection to RabbitMQ")

	return true
}

// Connect establishes a connection to rabbitMQ.
func Connect() {
	for i := 0; i < 3; i++ {
		if status := connect(); status {
			break
		}

		time.Sleep(20 * time.Second)
	}
}

// GetChannel returns RabbitMQ channel.
func GetChannel() *amqp.Channel {
	return chann
}
