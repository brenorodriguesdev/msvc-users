package models

import "time"

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	CompanyEmail *string   `json:"company_email,omitempty"`
	Password     string    `json:"password"`
	Name         string    `json:"name"`
	DateOfBirth  time.Time `json:"date_of_birth"`
}
