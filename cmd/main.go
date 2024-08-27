package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/the-jey/gomushroomapi/controllers"
)

func main() {
	r := mux.NewRouter()

	// Hello, World route!
	r.HandleFunc("/", controllers.Home).Methods("GET")

	// Start server
	fmt.Println("Server is running: 127.0.0.1:8080 🏃")
	log.Fatal(http.ListenAndServe(":8080", r))
}
