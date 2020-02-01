package utils

import (
	"encoding/json"
	"net/http"

	"github.com/jz222/rest-api-with-job-queue/api/models"
)

// RespondWithSuccess responds with a formatted success message.
func RespondWithSuccess(w http.ResponseWriter, payload interface{}) {
	response := models.SuccessResponse{
		Code:    http.StatusOK,
		Ok:      true,
		Payload: payload,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RespondWithError responds with a formatted error message.
func RespondWithError(w http.ResponseWriter, code int, message string) {
	response := models.ErrorResponse{
		Code:    code,
		Ok:      false,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
