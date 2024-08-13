package main

import (
	"log"
	"net/http"

	"github.com/Akshat120/Task-Management-System/api"
	"github.com/Akshat120/Task-Management-System/internal/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()
	defer db.DB.Close()

	r := mux.NewRouter()

	r.HandleFunc("/health", api.HealthCheckHandler)

	r.HandleFunc("/", api.HealthCheckHandler)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
