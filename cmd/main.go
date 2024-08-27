package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/the-jey/gomushroomapi/controllers"
	"github.com/the-jey/gomushroomapi/db"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	// Loading ENV variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found ❌")
	}
}

func main() {
	// Connected to the database
	MongoClient := db.NewMongoClient()
	defer db.DisconnectMongoClient(MongoClient)

	// Testing the database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := MongoClient.Ping(ctx, readpref.Primary()); err != nil {
		// panic("Error while pinging the database ❌")
		panic(err)
	} else {
		fmt.Println("Database successfully connected ✅")
	}

	// Create new router
	r := mux.NewRouter()

	// Hello, World route!
	r.HandleFunc("/", controllers.Home).Methods("GET")

	// Start server
	fmt.Println("Server is running: 127.0.0.1:8080 🏃")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error when starting the server ❌")
	}
}
