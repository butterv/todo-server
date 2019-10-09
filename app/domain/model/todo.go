package model

import (
	"time"
)

// TodoID is type of id in Todo.
type TodoID string

// Todo is a structure with task information.
type Todo struct {
	id            TodoID
	userID        UserID
	title         string
	project       *string
	label         *string
	description   *string
	scheduledDate *time.Time
	completed     bool
	createdAt     time.Time
	updatedAt     time.Time
	deletedAt     *time.Time
}

// NewTodo initializes Todo.
func NewTodo(userID UserID, title string, project, label, description *string, scheduledDate *time.Time) *Todo {
	return &Todo{
		userID:        userID,
		title:         title,
		project:       project,
		label:         label,
		description:   description,
		scheduledDate: scheduledDate,
	}
}

// GetID gets a id in Todo.
func (t *Todo) GetID() TodoID {
	return t.id
}

// GetUserID gets a user id in Todo.
func (t *Todo) GetUserID() UserID {
	return t.userID
}

// GetTitle gets a title in Todo.
func (t *Todo) GetTitle() string {
	return t.title
}

// GetLabel gets a label in Todo.
func (t *Todo) GetLabel() *string {
	return t.label
}

// GetDescription gets a description in Todo.
func (t *Todo) GetDescription() *string {
	return t.description
}

// GetScheduledDate gets a scheduled date in Todo.
func (t *Todo) GetScheduledDate() *time.Time {
	return t.scheduledDate
}

// Completed returns todo status.
func (t *Todo) Completed() bool {
	return t.completed
}
