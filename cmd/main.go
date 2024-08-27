package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Server is running: 127.0.0.1:8080 🏃")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
