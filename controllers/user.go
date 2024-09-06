package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/the-jey/gomushroomapi/errors"
	"github.com/the-jey/gomushroomapi/models"
	"github.com/the-jey/gomushroomapi/services"
	"github.com/the-jey/gomushroomapi/utils"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Get data from the body
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		errors.SendJSONErrorResponse(w, "Error parsing JSON data ❌", http.StatusBadRequest)
		return
	}

	// TODO: hash the password with bcrypt

	// Create the user
	id, s := services.NewUser(u)
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusInternalServerError)
		return
	}

	// Get the user by ID
	newU, s := services.GetUserByID(id)
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusInternalServerError)
		return
	}

	// Response
	utils.SendHttpJSONResponse(w, http.StatusCreated, newU)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

}
