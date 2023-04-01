package persistence

import (
	"fmt"

	"github.com/katsu105/hexagonal-todo/domain/entities"
	"github.com/katsu105/hexagonal-todo/domain/repository"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) repository.TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Save(task *entities.Task) error {
	// result := r.db.Create(task)
	fmt.Println("infrastructure/ Save!")
	return nil
}
