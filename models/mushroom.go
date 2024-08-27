package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Strenght string

const (
	Weak       Strenght = "Weak"
	Normal     Strenght = "Normal"
	Strong     Strenght = "Strong"
	Delusional Strenght = "Delusional"
)

type Mushroom struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	Origin    string             `bson:"origin,omitempty" json:"origin"`
	Price     float64            `bson:"price,omitempty" json:"price"`
	Quantity  int64              `bson:"quantity,omitempty" json:"quantity"`
	Strenght  Strenght           `bson:"strenght,omitempty" json:"strenght"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
