package infra

import (
	"msvc-users/domain/models"
	database "msvc-users/main/config"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Save(user models.User) error {
	result := database.DB.Create(&user)
	return result.Error
}

func (r *UserRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	result := database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
