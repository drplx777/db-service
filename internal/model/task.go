package model

import (
	"time"
)

type Task struct {
	ID          string     `db:"id" json:"id"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DoneAt      *time.Time `db:"done_at" json:"done_at,omitempty"`
	Completed   bool       `db:"completed" json:"completed"`
}
