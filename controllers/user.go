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

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Get data from the body
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		errors.SendJSONErrorResponse(w, "Error parsing JSON data ❌", http.StatusBadRequest)
		return
	}

	// Validate the user fields
	s := validation.CreateUserValidation(&u)
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusInternalServerError)
		return
	}

	// Hash the password with bcrypt
	hPass, err := utils.HashPassword(u.Password)
	if err != nil {
		errors.SendJSONErrorResponse(w, "Error hashing the password ❌", http.StatusInternalServerError)
		return
	}
	u.Password = hPass

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
	// TODO: with JWT token implementation
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
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

	u, s := services.GetUserByID(mongoID)
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusInternalServerError)
		return
	}

	utils.SendHttpJSONResponse(w, http.StatusOK, u)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allU, s := services.GetAllUsers()
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusInternalServerError)
		return
	}

	utils.SendHttpJSONResponse(w, http.StatusOK, allU)
}

func DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	_, s := services.DeleteAllUsers()
	if s != "" {
		errors.SendJSONErrorResponse(w, s, http.StatusInternalServerError)
		return
	}

	utils.SendHttpJSONResponse(w, http.StatusNoContent, nil)
}
