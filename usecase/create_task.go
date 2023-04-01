package usecase

import (
	"fmt"
	"time"

	"github.com/katsu105/hexagonal-todo/domain/entities"
	"github.com/katsu105/hexagonal-todo/domain/repository"
)

type CreateTaskInput struct {
	ID          int
	Title       string
	Description string
	Complate    bool
	DeadLine    *time.Time
}

type CreateTaskOutput struct {
	Task *entities.Task
}

type TaskCreator interface {
	Create(input *CreateTaskInput) (*CreateTaskOutput, error)
}

type createTask struct {
	taskRepository repository.TaskRepository
}

func NewCreateTask(taskRepository repository.TaskRepository) TaskCreator {
	return &createTask{taskRepository: taskRepository}
}

func (c *createTask) Create(input *CreateTaskInput) (*CreateTaskOutput, error) {
	fmt.Println("usecase/create_task Create! start")
	task, err := entities.NewTask(input.ID, input.Title, input.DeadLine)
	if err != nil {
		return nil, err
	}

	err = c.taskRepository.Save(task)
	if err != nil {
		return nil, err
	}

	fmt.Println("usecase/create_task Create! end")
	return &CreateTaskOutput{Task: task}, nil
}
