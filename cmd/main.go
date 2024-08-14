package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/Akshat120/Task-Management-System/api"
	"github.com/Akshat120/Task-Management-System/internal/db"
	"github.com/Akshat120/Task-Management-System/internal/handler"
	"github.com/Akshat120/Task-Management-System/internal/postgres"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbconn := db.InitDB()
	defer dbconn.Close()

	taskRepo := postgres.NewTaskRepo(dbconn)

	tmpl := template.Must(template.ParseFiles("templates/task.html"))
	taskHandler := handler.NewTaskHandler(tmpl, taskRepo)

	r := mux.NewRouter()

	r.HandleFunc("/health", api.HealthCheckHandler)

	r.HandleFunc("/", api.HealthCheckHandler)

	r.HandleFunc("/task", taskHandler.Handle)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
