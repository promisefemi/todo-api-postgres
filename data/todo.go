package data

import (
	"database/sql"
	"fmt"
	"time"
)

type Todo struct {
	ID          int          `json:"id"`
	Text        string       `json:"text"`
	Description string       `json:"description"`
	Completed   bool         `json:"completed"`
	CompletedAt sql.NullTime `json:"completed_at"`
	CreatedAt   time.Time    `json:"created_at"`
}

func (t *Todo) Validate() (bool, error) {

	if t.Text == "" {
		return false, fmt.Errorf("text field is required")
	}
	if t.Description == "" {
		return false, fmt.Errorf("description field is required")
	}

	return true, nil
}
