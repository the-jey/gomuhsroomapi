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

func GetAllUsers() ([]*models.User, string) {
	// Get User collection
	col := db.GetUsersCollection()

	// Create & defer context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var uArray []*models.User

	// Getting a cursor from mongoDB
	cur, err := col.Find(ctx, bson.M{})
	if err != nil {
		return uArray, "Error getting all the users ❌"
	}

	// Decode the cursor & append uArray
	for cur.Next(ctx) {
		var u *models.User
		if err := cur.Decode(&u); err != nil {
			return uArray, "Error during the decoding of all the users ❌"
		}
		uArray = append(uArray, u)
	}

	// Error during decoding & close the cursor
	if err := cur.Err(); err != nil {
		return uArray, "Error with the mongoDB cursor of all the users ❌"
	}
	cur.Close(ctx)

	return uArray, ""
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

func GetUserByUsername(uname string) (*models.User, string) {
	// Get User collection
	col := db.GetUsersCollection()

	// Create & defer context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a filter
	filter := bson.M{"username": uname}

	var u *models.User
	if err := col.FindOne(ctx, filter).Decode(&u); err != nil {
		return u, "Error getting the user by ID ❌"
	}

	return u, ""
}

func GetUserByEmail(email string) (*models.User, string) {
	// Get User collection
	col := db.GetUsersCollection()

	// Create & defer context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a filter
	filter := bson.M{"email": email}

	var u *models.User
	if err := col.FindOne(ctx, filter).Decode(&u); err != nil {
		return u, "Error getting the user by email ❌"
	}

	return u, ""
}

func DeleteAllUsers() (int64, string) {
	// Get User collection
	col := db.GetUsersCollection()

	// Create & defer context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Delete all the documents
	res, err := col.DeleteMany(ctx, bson.M{})
	if err != nil {
		return 0, "Error deleting all the users ❌"
	}

	// Check if count == 0
	count := res.DeletedCount
	if count == 0 {
		return count, "Error deleting all the users ❌"
	}

	// Return number of document deleted
	return res.DeletedCount, ""
}

func DeleteUserByID(id primitive.ObjectID) (int64, string) {
	// Get User collection
	col := db.GetUsersCollection()

	// Create & defer the context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a filter
	filter := bson.M{"_id": id}

	// Delete one document
	res, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return 0, "Error deleting the user by ID"
	}

	// Check if count == 0
	count := res.DeletedCount
	if count == 0 {
		return count, "Error deleting the user by ID"
	}

	// Return number of document deleted
	return res.DeletedCount, ""
}
