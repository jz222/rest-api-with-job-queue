package main

import (
	"log"
	"net/http"

	"github.com/jz222/rest-api-with-job-queue/worker/config"
	"github.com/jz222/rest-api-with-job-queue/worker/keys"
	"github.com/jz222/rest-api-with-job-queue/worker/routes"
	"github.com/jz222/rest-api-with-job-queue/worker/services/queue"
	"github.com/jz222/rest-api-with-job-queue/worker/services/worker"
)

func init() {
	keys.Load()
	queue.Connect()
}

func main() {
	router := routes.GetRouter()

	server := &http.Server{
		Addr:         ":" + keys.Get().Port,
		Handler:      router,
		ReadTimeout:  config.Server.ReadTimeout,
		WriteTimeout: config.Server.WriteTimeout,
		IdleTimeout:  config.Server.IdleTimeout,
	}

	go worker.Start()

	log.Fatal(server.ListenAndServe())
}
