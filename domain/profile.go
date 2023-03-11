package domain

import "context"

type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID int64) (*Profile, error)
}
