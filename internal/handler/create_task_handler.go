package handler

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/Akshat120/Task-Management-System/internal/repos"
)

type createTaskHandler struct {
	tmpl *template.Template
	repo repos.TaskRepo
}

func NewCreateTaskHandler(tmpl *template.Template, repo repos.TaskRepo) *createTaskHandler {
	return &createTaskHandler{repo: repo, tmpl: tmpl}
}

func (handler *createTaskHandler) HandleShowForm(w http.ResponseWriter, r *http.Request) {
	if err := handler.tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func (handler *createTaskHandler) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var due_date time.Time
	var err error
	layout := "2006-01-02"
	if r.FormValue("due_date") != "" {
		due_date, err = time.Parse(layout, r.FormValue("due_date"))
		if err != nil {
			fmt.Println("err:", err)
			http.Error(w, "Error in parsing due_date", http.StatusInternalServerError)
			return
		}
	}

	task := repos.Task{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Status:      r.FormValue("status"),
		DueDate:     &due_date,
	}

	drn_id, err := handler.repo.Upsert(&task)
	if err != nil {
		fmt.Println("err:", err)
		http.Error(w, "Error in creating task", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/preview?id=%v", drn_id.String()), http.StatusMovedPermanently)
}
