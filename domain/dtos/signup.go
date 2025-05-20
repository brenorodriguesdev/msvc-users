package dtos

import "time"

type SignupInput struct {
	Email       string
	Name        string
	DateOfBirth time.Time
}
