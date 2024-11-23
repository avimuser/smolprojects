package task

import "time"

type TaskStatus int

const (
	Incomplete TaskStatus = iota
	Started    TaskStatus = iota
	Completed  TaskStatus = iota
)

type Task struct {
	Desc      string
	Status    TaskStatus
	CreateAt  time.Time
	UpdatedAt time.Time
}
