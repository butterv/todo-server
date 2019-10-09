package model

import "time"

type TodoID string

type Todo struct {
	id            TodoID
	title         string
	project       string
	label         string
	description   string
	scheduledDate time.Time
	createdAt     time.Time
	updatedAt     time.Time
	deletedAt     *time.Time
}

func (t *Todo) GetID() TodoID {
	return t.id
}

func (t *Todo) GetTitle() string {
	return t.title
}

func (t *Todo) GetLabel() string {
	return t.label
}

func (t *Todo) GetDescription() string {
	return t.description
}

func (t *Todo) GetScheduledDate() time.Time {
	return t.scheduledDate
}
