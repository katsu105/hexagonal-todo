package repository

import (
	"github.com/katsu105/hexagonal-todo/domain/entities"
)

// TaskRepository は、タスクエンティティの永続化操作を抽象化したインターフェイス
type TaskRepository interface {
	// FindAll は、すべてのタスクを取得します。
	// FindAll() ([]*entities.Task, error)

	// FindByID は、指定されたIDのタスクを取得します。
	// FindByID(id string) (*entities.Task, error)

	// Save は、タスクを保存します。
	Save(task *entities.Task) error

	// Update は、タスクを更新します。
	// Update(task *entities.Task) error

	// Delete は、指定されたIDのタスクを削除します。
	// Delete(id string) error
}
