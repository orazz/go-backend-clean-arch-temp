package domain

import (
	"context"
	"time"
)

const (
	CollectionTask = "tasks"
)

type Task struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" form:"title" binding:"required" json:"title"`
	UserID    int64     `json:"userID" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID int64) ([]Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID int64) ([]Task, error)
}
