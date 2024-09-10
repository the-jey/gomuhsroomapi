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
	"github.com/the-jey/gomushroomapi/middlewares"
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

	// Mushrooms routes
	r.HandleFunc("/mushrooms", controllers.GetAllMushrooms).Methods("GET")
	r.HandleFunc("/mushrooms", controllers.DeleteAllMushrooms).Methods("DELETE")
	r.HandleFunc("/mushroom", controllers.CreateMushroom).Methods("POST")
	r.HandleFunc("/mushroom/{id}", controllers.GetOneMushroomByID).Methods("GET")
	r.HandleFunc("/mushroom/{id}", controllers.UpdateMushroomByID).Methods("PUT")
	r.HandleFunc("/mushroom/{id}", controllers.DeleteOneMushroomByID).Methods("DELETE")

	// Auth routes
	r.HandleFunc("/user/new", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/user/login", controllers.LoginUser).Methods("POST")

	// Users routes
	// r.Handle("/users", middlewares.IsLogin(http.HandlerFunc(controllers.GetAllUsers))).Methods("GET")
	r.HandleFunc("/users", middlewares.IsLogin(controllers.GetAllUsers)).Methods("GET")
	r.HandleFunc("/users", controllers.DeleteAllUsers).Methods("DELETE")
	r.HandleFunc("/user/{id}", controllers.GetUserByID).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.DeleteUserByID).Methods("DELETE")
	r.HandleFunc("/user/username/{username}", controllers.GetUserByUsername).Methods("GET")
	r.HandleFunc("/user/email/{email}", controllers.GetUserByEmail).Methods("GET")

	// Start server
	fmt.Println("Server is running: 127.0.0.1:8080 🏃")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error when starting the server ❌")
	}
}
