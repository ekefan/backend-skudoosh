package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Message struct {
	Text string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Open and connect to the SQLite database
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	defer db.Close()

	// Create table if not exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS messages (id INTEGER PRIMARY KEY AUTOINCREMENT, text TEXT)")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	// Insert a message into the database
	_, err = db.Exec("INSERT INTO messages (text) VALUES (?)", "Hello, Database!")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	// Query the message from the database
	var text string
	row := db.QueryRow("SELECT text FROM messages WHERE id = ?", 1)
	err = row.Scan(&text)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	// Create a Message struct instance
	message := Message{
		Text: text,
	}

	// Marshal the Message struct to JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
