package repository

import (
	"time"

	"github.com/istsh/todo-sample/app/domain/model"
)

// TodoRepositoryRefer is a readonly repository for todos.
type TodoRepositoryRefer interface {
	FindByID(id model.TodoID) (*model.Todo, error)
	FindByProject(project string) ([]*model.Todo, error)
	FindByLabel(label string) ([]*model.Todo, error)
	FindByScheduledDate(scheduledDate time.Time) ([]*model.Todo, error)
}

// TodoRepositoryModify is a read/write repository for todos.
type TodoRepositoryModify interface {
	TodoRepositoryRefer

	Create(id model.TodoID, title string, project, label, description *string, scheduledDate *time.Time) (*model.Todo, error)
	UnsafeUpdate(todo *model.Todo) error
	DeleteByID(id model.TodoID) error
}
