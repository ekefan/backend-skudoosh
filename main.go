package main

import (
    "fmt"
    "net/http"
	"encoding/json"
    "os"
)

type Message struct {
    Text string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Create a Message struct instance
    message := Message{
        Text: "Hello, World!",
    }

    // Marshal the Message struct to JSON
    jsonData, err := json.Marshal(message)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Set Content-Type header to application/json
    w.Header().Set("Content-Type", "application/json")

    // Write the JSON response
    _, err = w.Write(jsonData)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
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