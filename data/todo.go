package data

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID          int       `json:"id"`
	Text        string    `json:"text"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CompletedAt sql.NullTime `json:"completed_at"`
	CreatedAt   time.Time `json:"created_at"`
}

func (t *Todo) Validate() (bool, error) {
	return true, nil
}
