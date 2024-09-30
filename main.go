package main

import (
    "fmt"
    "net/http"
)

// handler function for the root URL
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World! Welcome to my simple Go web app.")
}

func main() {
    // Register the handler function
    http.HandleFunc("/", homeHandler)

    // Start the server on port 8080
    fmt.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}
