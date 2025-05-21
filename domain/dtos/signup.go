package dtos

import "time"

type SignupInput struct {
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
}
