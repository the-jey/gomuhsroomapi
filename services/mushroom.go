package services

import (
	"context"
	"time"

	"github.com/the-jey/gomushroomapi/db"
	"github.com/the-jey/gomushroomapi/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewMushroom(m models.Mushroom) (primitive.ObjectID, error) {
	col := db.GetMushroomsCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	out, err := col.InsertOne(ctx, m)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return out.InsertedID.(primitive.ObjectID), nil
}
