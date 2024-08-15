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
	// TODO: use a config struct to maintain env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbconn := db.InitDB()
	defer dbconn.Close()

	taskRepo := postgres.NewTaskRepo(dbconn)

	previewTaskTemplate := template.Must(template.ParseFiles("templates/preview_task.html"))
	previewTaskHandler := handler.NewPreviewTaskHandler(previewTaskTemplate, taskRepo)

	createTaskTemplate := template.Must(template.ParseFiles("templates/new_task.html"))
	createTaskHandler := handler.NewCreateTaskHandler(createTaskTemplate, taskRepo)

	r := mux.NewRouter()

	r.HandleFunc("/health", api.HealthCheckHandler)

	r.HandleFunc("/", api.HealthCheckHandler)

	r.HandleFunc("/preview", previewTaskHandler.Handle).Methods("GET")
	r.HandleFunc("/new", createTaskHandler.HandleShowForm).Methods("GET")
	r.HandleFunc("/new", createTaskHandler.HandleCreateTask).Methods("POST")
	// TODO: list_task, update and delete
	// r.HandleFunc("/update", ).Methods("POST")
	// r.HandleFunc("/delete", ).Methods("POST")

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
