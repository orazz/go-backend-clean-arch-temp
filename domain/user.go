package domain

import (
	"context"
	"time"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	FetchAll(c context.Context) ([]*User, error)
	GetByEmail(c context.Context, email string) (*User, error)
	GetByID(c context.Context, id int64) (*User, error)
}
