package usecases

import "msvc-users/domain/dtos"

type SignIn interface {
	SignIn(input dtos.SignInInput) (string, error)
}
