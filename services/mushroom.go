package services

import (
	"context"
	"time"

	"github.com/the-jey/gomushroomapi/db"
	"github.com/the-jey/gomushroomapi/models"
	"github.com/the-jey/gomushroomapi/validation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewMushroom(m models.Mushroom) (primitive.ObjectID, string) {
	// Mushroom model validation
	s := validation.CreateMushroomValidation(&m)
	if s != "" {
		return primitive.NilObjectID, s
	}

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
