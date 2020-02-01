package worker

import (
	"log"

	"github.com/jz222/rest-api-with-job-queue/worker/config"
	"github.com/jz222/rest-api-with-job-queue/worker/services/queue"
)

// Start starts to consume messages from the queue.
func Start() {
	channel := queue.GetChannel()

	msgs, err := channel.Consume(config.Queue.QueueName, "", false, false, false, false, nil)
	if err != nil {
		log.Println("Failed to consume with error:" + err.Error())
	}

	for msg := range msgs {
		log.Printf("[Worker] Received a message: %s", msg.Body)
		log.Print("[Worker] Working...")
		log.Print("[Worker] Finished work")
		msg.Ack(false)
	}
}
