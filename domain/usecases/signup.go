package usecases

import "msvc-users/domain/dtos"

type SignUp interface {
	SignUp(input dtos.SignupInput) error
}
