package entites

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	ID       int        `json:"id"`
	Title    string     `json:"title"`
	Complete bool       `json:"complete"`
	DueDate  *time.Time `json:"duedate"`
}

func NewTask(id int, title string, duedate *time.Time) (*Task, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	now := time.Now()
	if duedate.Before(now) {
		return nil, errors.New("invalid deadline")
	}

	return &Task{
		ID:      id,
		Title:   title,
		DueDate: duedate,
	}, nil
}

func (t *Task) CompleteTask() error {
	// 完了するための処理
	fmt.Println("domain/ complate!")
	return nil
}

func (t *Task) UpdateTask(title string, duedate *time.Time) error {
	// updateするための処理
	fmt.Println("domain/ update!")
	return nil
}
