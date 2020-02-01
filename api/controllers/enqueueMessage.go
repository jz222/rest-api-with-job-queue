package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jz222/rest-api-with-job-queue/api/models"
	"github.com/jz222/rest-api-with-job-queue/api/services/queue"
	"github.com/jz222/rest-api-with-job-queue/api/utils"
)

// EnqueueMessage enqueues the message sent in the request body.
func EnqueueMessage(w http.ResponseWriter, r *http.Request) {
	var msg models.Message

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&msg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = queue.EnqueueMessage([]byte(msg.Message))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.RespondWithSuccess(w, "Message successfully enqueued.")
	}
}
