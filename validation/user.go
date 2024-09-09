package validation

import (
	"context"
	"net/http"
	"net/mail"
	"time"

	"github.com/the-jey/gomushroomapi/db"
	"github.com/the-jey/gomushroomapi/models"
	"github.com/the-jey/gomushroomapi/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUserValidation(u *models.User) string {
	// 'Username' field
	if u.Username == "" {
		s := "Please put an 'username' field ❌"
		return s
	}
	if CheckUsernameExist(u.Username) {
		s := "'username' already exist ❌"
		return s
	}
	if (len(u.Username) < 3) || (len(u.Username) > 224) {
		s := "'Username' must be between 3 and 224 characters ❌"
		return s
	}

	// 'Email' field
	if u.Email == "" {
		s := "Please put an 'email' field ❌"
		return s
	}
	if !IsValidEmail(u.Email) {
		s := "Please put a valid email ❌"
		return s
	}
	if CheckUserEmailExist(u.Email) {
		s := "'email' already exist ❌"
		return s
	}
	if (len(u.Email) < 3) || (len(u.Email) > 224) {
		s := "User 'email' must be between 3 and 224 characters ❌"
		return s
	}

	// 'Password' field
	if u.Password == "" {
		s := "Please put a 'password' field ❌"
		return s
	}
	if (len(u.Password) < 3) || (len(u.Password) > 224) {
		s := "'Password' must be between 3 and 224 characters ❌"
		return s
	}

	return ""
}

func CheckUserEmailExist(e string) bool {
	// Get the user collection
	col := db.GetUsersCollection()

	// Create & defer context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a filter
	filter := bson.M{"email": e}

	// Return an error if document doesn't exist
	err := col.FindOne(ctx, filter).Err()

	return err != mongo.ErrNoDocuments
}

func CheckUsernameExist(u string) bool {
	// Get the user collection
	col := db.GetUsersCollection()

	// Create & defer context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a filter
	filter := bson.M{"username": u}

	// Return an error if document doesn't exist
	err := col.FindOne(ctx, filter).Err()

	return err != mongo.ErrNoDocuments
}

func IsValidEmail(e string) bool {
	_, err := mail.ParseAddress(e)
	return err == nil
}

func IsValidUsername(u string) string {
	if u == "" {
		s := "Please put an 'username' field ❌"
		return s
	}
	if !CheckUsernameExist(u) {
		s := "'username' doesn't exist ❌"
		return s
	}
	if (len(u) < 3) || (len(u) > 224) {
		s := "'Username' must be between 3 and 224 characters ❌"
		return s
	}

	return ""
}

func LoginValidationPayload(lp *models.LoginPayload) (primitive.ObjectID, string, int) {
	if lp.Email == "" && lp.Username == "" {
		return primitive.NilObjectID, "'email' or 'username' invalid ❌", http.StatusBadRequest
	}

	// Login with email
	if lp.Email != "" {
		// Check if email is valid
		if !IsValidEmail(lp.Email) {
			return primitive.NilObjectID, "'email' is invalid ❌", http.StatusBadRequest
		}

		// Get the user by 'email' field and return the ID
		u, s := services.GetUserByEmail(lp.Email)
		if s != "" {
			return primitive.NilObjectID, s, http.StatusBadRequest
		}

		return u.ID, "", http.StatusOK
	}

	// Login with username here
	if lp.Username != "" {

		// Check if the username is valid
		s := IsValidUsername(lp.Username)
		if s != "" {
			return primitive.NilObjectID, s, http.StatusBadRequest
		}

		// Get the user by Username
		u, s := services.GetUserByUsername(lp.Username)
		if s != "" {
			return primitive.NilObjectID, s, http.StatusBadRequest
		}

		return u.ID, "", http.StatusOK
	}

	// TODO: validation of the payload 'password' field

	return primitive.NilObjectID, "login action failed ❌", http.StatusInternalServerError
}
