package contracts

import "msvc-users/domain/models"

type UserRepository interface {
	Save(user models.User) error
	FindByEmail(email string) (models.User, error)
}
