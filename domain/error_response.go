package domain

import "errors"

type ErrorResponse struct {
	Message string `json:"message"`
}

var (
	ErrNotFound = errors.New("Resquesed item is not found!")
)
