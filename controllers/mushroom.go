package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/the-jey/gomushroomapi/errors"
	"github.com/the-jey/gomushroomapi/models"
	"github.com/the-jey/gomushroomapi/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetAllMushrooms(w http.ResponseWriter, r *http.Request) {
	allM, s := services.GetAllMushrooms()
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allM)
}

func GetOneMushroomByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, ok := mux.Vars(r)["id"]
	if (!ok) || (id == "") {
		errors.SendJSONErrorResponse(w, "Please give an 'id' parameter ❌", http.StatusBadRequest)
		return
	}

	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		errors.SendJSONErrorResponse(w, "'id' parameter is not valid ❌", http.StatusBadRequest)
		return
	}

	m, s := services.GetMushroomByID(mongoID)
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}
