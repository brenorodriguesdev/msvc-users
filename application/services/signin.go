package services

import (
	"errors"
	"msvc-users/application/contracts"
	"msvc-users/domain/dtos"
	"msvc-users/domain/usecases"
)

type SignInService struct {
	repo          contracts.UserRepository
	hasherCompare contracts.HasherCompare
	encrypter     contracts.Encrypter
}

func NewSignInService(r contracts.UserRepository, h contracts.HasherCompare, e contracts.Encrypter) *SignInService {
	return &SignInService{repo: r, hasherCompare: h, encrypter: e}
}

var _ usecases.SignIn = (*SignInService)(nil)

func (s *SignInService) SignIn(input dtos.SignInInput) (string, error) {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New("e-mail or password is invalid")
	}

	hashedPassword, err := s.hasherCompare.Compare(input.Password, user.Password)
	if err != nil || !hashedPassword {
		return "", errors.New("e-mail or password is invalid")
	}

	token, err := s.encrypter.Encrypt(user)
	if err != nil {
		return "", errors.New("e-mail or password is invalid")
	}

	return token, nil
}
