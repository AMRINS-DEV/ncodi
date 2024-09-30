package main

import (
    "fmt"
    "net/http"
    "os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World! Welcome to my Go web app.")
}

func main() {
    http.HandleFunc("/", homeHandler)

    // Get the port from environment variable or default to 8080
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Printf("Starting server on :%s...\n", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}
