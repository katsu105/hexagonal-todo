package repository

import (
	"github.com/katsu105/hexagonal-todo/domain/entities"
)

// TaskRepository は、タスクエンティティの永続化操作を抽象化したインターフェイス
type TaskRepository interface {
	// FindAll 全てのタスクを取得
	// FindAll() ([]*entities.Task, error)

	// FindByID 指定されたIDのタスクを取得
	// FindByID(id string) (*entities.Task, error)

	// Save タスクを保存
	Save(task *entities.Task) error

	// Update タスクを更新
	// Update(task *entities.Task) error

	// Delete タスクを削除
	// Delete(id string) error
}
