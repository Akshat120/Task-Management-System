package repos

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	DrnID       uuid.UUID
	Title       string
	Description string
	Status      string
	DueDate     *time.Time
}

type TaskRepo interface {
	FindByDrnID(uuid.UUID) (*Task, error)
	Upsert(*Task) (uuid.UUID, error)
	Delete(uuid.UUID) (bool, error)
}
