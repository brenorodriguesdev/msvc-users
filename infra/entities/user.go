package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint   `json:"id" gorm:"primaryKey"`
	Email        string `json:"email" gorm:"unique;not null"`
	CompanyEmail string `json:"company_email"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	DateOfBirth  string `json:"date_of_birth"`
}
