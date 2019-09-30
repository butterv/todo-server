package model

import "time"

type TodoID string

type Todo struct {
	ID            TodoID
	Title         string
	Project       string
	Label         string
	Description   string
	ScheduledDate time.Time
}
