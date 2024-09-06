package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient() *mongo.Client {
	uri := os.Getenv("DB_URI")
	if uri == "" {
		log.Fatal("Set your 'DB_URI' environment variable ❌")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic("Error when creating a MongoDB connection ❌")
	}

	return client
}

func GetDatabase() *mongo.Database {
	return NewMongoClient().Database("gomushroomapi")
}

func GetMushroomsCollection() *mongo.Collection {
	return GetDatabase().Collection("muhsrooms")
}

func GetUsersCollection() *mongo.Collection {
	return GetDatabase().Collection("users")
}

func DisconnectMongoClient(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic("Error when disconnecting MongoDB connection ❎")
	}
}
