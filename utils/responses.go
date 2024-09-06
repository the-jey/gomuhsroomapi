package utils

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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

func HashPassword(p string) (string, error) {
	bHasedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bHasedPassword), nil
}
