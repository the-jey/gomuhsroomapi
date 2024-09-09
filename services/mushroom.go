package services

import (
	"context"
	"time"

	"github.com/the-jey/gomushroomapi/db"
	"github.com/the-jey/gomushroomapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewMushroom(m models.Mushroom) (primitive.ObjectID, string) {
	// CreatedAt & UpdatedAt field
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	// Get Mushroom collection
	col := db.GetMushroomsCollection()

	// Create & defer context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the mushroom in the database
	out, err := col.InsertOne(ctx, m)
	if err != nil {
		return primitive.NilObjectID, "Error creating the mushroom in the database ❌"
	}

	return out.InsertedID.(primitive.ObjectID), ""
}

func GetMushroomByID(id primitive.ObjectID) (*models.Mushroom, string) {
	// Get Mushroom collection
	col := db.GetMushroomsCollection()

	// Create & defer context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a filter
	filter := bson.M{"_id": id}

	var m *models.Mushroom
	if err := col.FindOne(ctx, filter).Decode(&m); err != nil {
		return m, "Error getting the mushroom by ID ❌"
	}

	return m, ""
}

func GetAllMushrooms() ([]*models.Mushroom, string) {
	// Get Mushroom collection
	col := db.GetMushroomsCollection()

	// Create & defer context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a mushroom array
	var mArray []*models.Mushroom

	// Getting a cursor from MongoDB
	cur, err := col.Find(ctx, bson.D{{}})
	if err != nil {
		return mArray, "Error getting all the mushrooms ❌"
	}

	// Decode the cursor & append to mArray
	for cur.Next(ctx) {
		var m *models.Mushroom
		if err := cur.Decode(&m); err != nil {
			return mArray, "Error during decoding all the mushrooms ❌"
		}
		mArray = append(mArray, m)
	}
	// Error during decoding & close the cursor
	if err := cur.Err(); err != nil {
		return mArray, "Error with the mongoDB cursor of all the mushrooms ❌"
	}
	cur.Close(ctx)

	return mArray, ""
}

func DeleteMushroomByID(id primitive.ObjectID) (int64, string) {
	// Get Mushroom collection
	col := db.GetMushroomsCollection()

	// Create & defer context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a filter
	filter := bson.M{"_id": id}

	// Delete one document
	res, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return 0, "Error deleting the mushroom by ID ❌"
	}

	// Check if count == 0
	count := res.DeletedCount
	if count == 0 {
		return count, "Error deleting the mushroom by ID ❌"
	}

	// Return number of document deleted
	return res.DeletedCount, ""
}

func DeleteAllMushrooms() (int64, string) {
	// Get Mushroom collection
	col := db.GetMushroomsCollection()

	// Create & defer context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Delete one document
	res, err := col.DeleteMany(ctx, bson.M{})
	if err != nil {
		return 0, "Error deleting all the mushrooms ❌"
	}

	// Check if count == 0
	count := res.DeletedCount
	if count == 0 {
		return count, "Error deleting all the mushrooms ❌"
	}

	// Return number of document deleted
	return res.DeletedCount, ""
}

func UpdateMushroomByID(id primitive.ObjectID, m models.Mushroom) string {
	// Update the time
	m.UpdatedAt = time.Now()

	// Get Mushroom collection
	col := db.GetMushroomsCollection()

	// Create & defer context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a filter
	filter := bson.M{"_id": id}

	// Create an update
	update := bson.M{"$set": m}

	// Update by ID
	_, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return "Error updating by id ❌"
	}

	return ""
}
