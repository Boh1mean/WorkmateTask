package model

import (
	"time"
)

type Task struct {
	ID        string        `json:"id"`
	Status    string        `json:"status"`
	Duration  time.Duration `json:"duration"`
	CreatedAt time.Time     `json:"createdat"`
}
