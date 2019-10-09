package usecase

import (
	"time"

	"github.com/istsh/todo-sample/app/domain/model"
)

// UserUsecase is an user usecase.
type TodoUsecase interface {
	CreateTodo(title string, project, label, description *string, scheduledDate *time.Time) (*model.Todo, error)
	GetIncompletTodosByUserID(userID model.UserID) ([]*model.Todo, error)
	GetIncompletTodosByProject(userID model.Todo, project string) ([]*model.Todo, error)
	GetIncompletTodosByLabel(userID model.Todo, label string) ([]*model.Todo, error)
	GetIncompletTodosByScheduledDate(userID model.Todo, scheduledDate time.Time) ([]*model.Todo, error)
	GetCompletedTodosByUserID(userID model.Todo) ([]*model.Todo, error)
	UpdateTitle(todoID model.TodoID, title string) error
	UpdateProject(todoID model.TodoID, project *string) error
	UpdateLabel(todoID model.TodoID, label *string) error
	UpdateDescription(todoID model.TodoID, description *string) error
	UpdateScheduledDate(todoID model.TodoID, scheduledDate *time.Time) error
	Complete(todoID model.TodoID) error
	Incomplete(todoID model.TodoID) error
	DeleteTodo(todoID model.TodoID) error
}
