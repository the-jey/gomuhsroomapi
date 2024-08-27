package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/the-jey/gomushroomapi/models"
	"github.com/the-jey/gomushroomapi/services"
)

func CreateMushroom(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var m models.Mushroom
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Println("Error parsing JSON data ❌")
		http.Error(w, "Error parsing JSON data ❌", http.StatusBadRequest)
		return
	}

	id, err := services.NewMushroom(m)
	if err != nil {
		log.Println("Error creating the mushroom in the database ❌")
		http.Error(w, "Error creating the mushroom in the database ❌", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Mushroom{ID: id})
}
