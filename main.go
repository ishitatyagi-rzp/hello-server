package main

import (
	"fmt"
	"net/http"
)

// 1. Handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// 2. Main function = entry point
func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
