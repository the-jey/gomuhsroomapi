package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorJSONPayload struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func SendJSONErrorResponse(w http.ResponseWriter, err string, status int) {
	log.Println(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorJSONPayload{
		Status:       status,
		ErrorMessage: err,
	})
}
