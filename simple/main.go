package main

import (
	"fmt"
	"net/http"
)

// Application entry point
func main() {
	// Define routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})
	// Run web server
	http.ListenAndServe(":8080", nil)
}
