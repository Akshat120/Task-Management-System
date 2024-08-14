package postgres

import (
	"fmt"
	"time"

	"github.com/Akshat120/Task-Management-System/internal/repos"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type taskPgRepo struct {
	DB *pg.DB
}

var _ repos.TaskRepo = taskPgRepo{}

type Task struct {
	DrnId       uuid.UUID  `pg:"drn_id,pk,type:uuid"`
	Title       string     `pg:"title"`
	Description string     `pg:"description"`
	Status      string     `pg:"status"`
	DueDate     *time.Time `pg:"due_date"`
}

func NewTaskRepo(db *pg.DB) repos.TaskRepo {
	return taskPgRepo{db}
}

// FindByDrnID implements repos.TaskRepo.
func (t taskPgRepo) FindByDrnID(drnId uuid.UUID) (*repos.Task, error) {
	task := repos.Task{}
	err := t.DB.Model(&task).
		Where("drn_id = ?", drnId).
		Select()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return &task, nil
}

// Delete implements repos.TaskRepo.
func (t taskPgRepo) Delete(uuid.UUID) (bool, error) {
	panic("unimplemented")
}

// Upsert implements repos.TaskRepo.
func (t taskPgRepo) Upsert(*repos.Task) (uuid.UUID, error) {
	panic("unimplemented")
}
