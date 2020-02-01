package main

import (
	"log"
	"net/http"

	"github.com/jz222/rest-api-with-job-queue/api/config"
	"github.com/jz222/rest-api-with-job-queue/api/keys"
	"github.com/jz222/rest-api-with-job-queue/api/routes"
	"github.com/jz222/rest-api-with-job-queue/api/services/queue"
)

func init() {
	keys.Load()
	queue.Connect()
}

func main() {
	server := &http.Server{
		Addr:         ":" + keys.Get().Port,
		Handler:      routes.GetRouter(),
		ReadTimeout:  config.Server.ReadTimeout,
		WriteTimeout: config.Server.WriteTimeout,
		IdleTimeout:  config.Server.IdleTimeout,
	}

	log.Fatal(server.ListenAndServe())
}
