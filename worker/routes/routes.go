package routes

import (
	"net/http"
)

var router = http.NewServeMux()

// GetRouter creates and returns a new router.
func GetRouter() *http.ServeMux {
	router.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {})

	return router
}
