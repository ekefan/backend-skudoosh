package main

import (
    "fmt"
    "net/http"
    "os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    // Retrieve port from environment variable, defaulting to 8080 if not set
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Start server with specified host and port
    host := os.Getenv("HOST")
    if host == "" {
        host = "0.0.0.0"
    }

    http.HandleFunc("/", helloHandler)
    fmt.Printf("Starting server on %s:%s\n", host, port)
    if err := http.ListenAndServe(host+":"+port, nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}