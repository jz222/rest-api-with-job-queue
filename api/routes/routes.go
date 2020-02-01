package routes

import (
	"github.com/jz222/rest-api-with-job-queue/api/controllers"
	"net/http"
)

var router = http.NewServeMux()

// GetRouter creates and returns a new router.
func GetRouter() *http.ServeMux {
	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {})
	router.HandleFunc("/addjob", controllers.EnqueueMessage)

	return router
}
