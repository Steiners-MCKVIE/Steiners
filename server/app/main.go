package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a request handler function.
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}

	// Register the handler function for a specific route pattern.
	http.HandleFunc("/", handler)

	// Specify the port to listen on.
	port := ":8080"

	// Start the HTTP server.
	fmt.Printf("Server is listening on port %s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

