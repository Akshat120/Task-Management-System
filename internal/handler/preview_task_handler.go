package handler

import (
	"net/http"
	"text/template"

	"github.com/Akshat120/Task-Management-System/internal/repos"
	"github.com/google/uuid"
)

type PreviewTaskHandler struct {
	repo repos.TaskRepo
	tmpl *template.Template
}

func NewTaskHandler(tmpl *template.Template, repo repos.TaskRepo) *PreviewTaskHandler {
	return &PreviewTaskHandler{tmpl: tmpl, repo: repo}
}

func (handler *PreviewTaskHandler) Handle(w http.ResponseWriter, r *http.Request) {

	drnIDStr := r.URL.Query().Get("id")
	drnID, err := uuid.Parse(drnIDStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := handler.repo.FindByDrnID(drnID)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// Render the task using the template
	if err := handler.tmpl.Execute(w, task); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
