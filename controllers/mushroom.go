package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/the-jey/gomushroomapi/errors"
	"github.com/the-jey/gomushroomapi/models"
	"github.com/the-jey/gomushroomapi/services"
	"github.com/the-jey/gomushroomapi/utils"
	"github.com/the-jey/gomushroomapi/validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMushroom(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var m models.Mushroom
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		errors.SendJSONErrorResponse(w, "Error parsing JSON data ❌", http.StatusBadRequest)
		return
	}

	// Mushroom model validation
	s := validation.CreateMushroomValidation(&m)
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusBadRequest)
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

	utils.SendHttpJSONResponse(w, http.StatusCreated, newM)
}

func GetAllMushrooms(w http.ResponseWriter, r *http.Request) {
	allM, s := services.GetAllMushrooms()
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusInternalServerError)
		return
	}

	utils.SendHttpJSONResponse(w, http.StatusOK, allM)
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

	utils.SendHttpJSONResponse(w, http.StatusOK, m)
}

func DeleteOneMushroomByID(w http.ResponseWriter, r *http.Request) {
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

	_, s := services.DeleteMushroomByID(mongoID)
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusInternalServerError)
		return
	}

	utils.SendHttpJSONResponse(w, http.StatusNoContent, nil)
}

func DeleteAllMushrooms(w http.ResponseWriter, r *http.Request) {
	_, s := services.DeleteAllMushrooms()
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusInternalServerError)
		return
	}

	utils.SendHttpJSONResponse(w, http.StatusNoContent, nil)
}

func UpdateMushroomByID(w http.ResponseWriter, r *http.Request) {
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

	var m models.Mushroom
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		errors.SendJSONErrorResponse(w, "Error parsing JSON data ❌", http.StatusBadRequest)
		return
	}

	// Validation the updateMushroom
	s := validation.UpdateMushroomValidation(&m)
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusBadRequest)
		return
	}

	s = services.UpdateMushroomByID(mongoID, m)
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusInternalServerError)
		return
	}

	updateM, s := services.GetMushroomByID(mongoID)
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusInternalServerError)
		return
	}

	utils.SendHttpJSONResponse(w, http.StatusOK, updateM)
}
