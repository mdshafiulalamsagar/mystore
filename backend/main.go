package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define a simple default route for testing the server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to myStore Backend API!")
	})

	// Set the server port
	port := ":8080"
	fmt.Printf("Server is starting on port %s\n", port)

	// Start the server and check for errors
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}