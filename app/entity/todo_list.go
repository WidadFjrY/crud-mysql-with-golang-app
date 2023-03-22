package entity

import (
	"time"
)

type TodoList struct {
	Id          int32
	Todo        string
	Description string
	IsFinished  bool
	CreatedAt   time.Time
}
