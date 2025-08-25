package main

import (
	"fmt"
	"net/http" // Go’s built-in HTTP package
)

// 1. Handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// w = "response writer" (you send data back to the client through this)
	// r = "request" (contains info about the client’s request)
	fmt.Fprintln(w, "Hello, World!")
}

// 2. Main function = entry point
func main() {
	// Register a route: when someone visits /hello → call helloHandler
	http.HandleFunc("/hello", helloHandler)

	// Start the server on port 8080
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
