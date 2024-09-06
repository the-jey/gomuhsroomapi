package validation

import (
	"context"
	"time"

	"github.com/the-jey/gomushroomapi/db"
	"github.com/the-jey/gomushroomapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUserValidation(u *models.User) {
	// TODO
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
