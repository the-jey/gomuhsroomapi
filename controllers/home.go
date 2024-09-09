package controllers

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	Message string `json:"message"`
	Creator string `json:"creator"`
}

// Hello, World handler
func Home(w http.ResponseWriter, r *http.Request) {
	data := Data{Message: "Hello, World!", Creator: "THE JEY"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
