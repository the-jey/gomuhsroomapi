package utils

import (
	"encoding/json"
	"net/http"
)

type HttpJSONResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func SendHttpJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	payload := HttpJSONResponse{
		Status: status,
		Data:   data,
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}
