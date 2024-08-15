package repos

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	DrnId       uuid.UUID  `json:"drn_id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Status      string     `json:"status"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

type TaskRepo interface {
	FindByDrnID(uuid.UUID) (*Task, error)
	Upsert(*Task) (uuid.UUID, error)
	Delete(uuid.UUID) (bool, error)
}
