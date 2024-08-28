package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/the-jey/gomushroomapi/errors"
	"github.com/the-jey/gomushroomapi/models"
	"github.com/the-jey/gomushroomapi/services"
)

func CreateMushroom(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var m models.Mushroom
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		errors.SendJSONErrorResponse(w, "Error parsing JSON data ❌", http.StatusBadRequest)
		return
	}

	id, s := services.NewMushroom(m)
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusInternalServerError)
		return
	}

	newM, s := services.GetMushroomByID(id)
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newM)
}
