package entities

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	ID       int        `json:"id"`
	Title    string     `json:"title"`
	Complete bool       `json:"complete"`
	DeadLine *time.Time `json:"deadline"`
}

func NewTask(id int, title string, deadline *time.Time) (*Task, error) {
	fmt.Println("domain/ NewTask! start")
	if title == "" {
		return nil, errors.New("title is required")
	}

	// now := time.Now()
	// if deadline.Before(now) {
	// 	return nil, errors.New("invalid deadline")
	// }

	fmt.Println("domain/ NewTask! end")

	return &Task{
		ID:       id,
		Title:    title,
		DeadLine: deadline,
	}, nil

}

func (t *Task) CompleteTask() error {
	// 完了するための処理
	fmt.Println("domain/ complate!")
	return nil
}

func (t *Task) UpdateTask(title string, deadline *time.Time) error {
	// updateするための処理
	fmt.Println("domain/ update!")
	return nil
}
