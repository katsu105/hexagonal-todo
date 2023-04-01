package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/katsu105/hexagonal-todo/usecase"
	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	createTask usecase.TaskCreator
}

func NewTaskHandler(createTask usecase.TaskCreator) *TaskHandler {
	return &TaskHandler{
		createTask: createTask,
	}
}

type createTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	// var req createTaskRequest
	fmt.Println("handlers/ CreateTask!")
	taskInput := usecase.CreateTaskInput{
		ID:       1,
		Title:    "todo title",
		DeadLine: &time.Time{},
	}
	task, err := h.createTask.Create(&taskInput)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, task)
}
