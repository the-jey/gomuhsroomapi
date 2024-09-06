package services

import (
	"context"
	"time"

	"github.com/the-jey/gomushroomapi/db"
	"github.com/the-jey/gomushroomapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewUser(u models.User) (primitive.ObjectID, string) {
	// Put the updatedAt and createxAt datetime
	u.UpdatedAt = time.Now()
	u.CreatedAt = time.Now()

	// Put the 'default' isAdmin value
	if !u.IsAdmin {
		u.IsAdmin = false
	}

	// Get the user collection
	col := db.GetUsersCollection()

	// Create & defer context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the user in the database
	out, err := col.InsertOne(ctx, u)
	if err != nil {
		return primitive.NilObjectID, "Error creating the mushroom in the database ❌"
	}

	// Return the id mongo object
	return out.InsertedID.(primitive.ObjectID), ""
}

func GetUserByID(id primitive.ObjectID) (*models.User, string) {
	// Get User collection
	col := db.GetUsersCollection()

	// Create & defer context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a filter
	filter := bson.M{"_id": id}

	var u *models.User
	if err := col.FindOne(ctx, filter).Decode(&u); err != nil {
		return u, "Error getting the user by ID ❌"
	}

	return u, ""
}
