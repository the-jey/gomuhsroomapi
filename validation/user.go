package validation

import (
	"context"
	"net/mail"
	"time"

	"github.com/the-jey/gomushroomapi/db"
	"github.com/the-jey/gomushroomapi/models"
	"go.mongodb.org/mongo-driver/bson"
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
		s := "Mushroom 'name' must be between 3 and 224 characters ❌"
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
		s := "Please put an 'password' field ❌"
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
