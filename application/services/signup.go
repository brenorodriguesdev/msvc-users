package services

import (
	"errors"
	"fmt"
	"msvc-users/application/contracts"
	"msvc-users/domain/dtos"
	"msvc-users/domain/models"
	"msvc-users/domain/usecases"
)

type SignUpService struct {
	repo              contracts.UserRepository
	hasher            contracts.Hasher
	passwordGenerator contracts.PasswordGenerator
	emailSender       contracts.EmailSender
}

func NewSignUpService(r contracts.UserRepository, h contracts.Hasher, p contracts.PasswordGenerator, e contracts.EmailSender) *SignUpService {
	return &SignUpService{repo: r, hasher: h, passwordGenerator: p, emailSender: e}
}

var _ usecases.SignUp = (*SignUpService)(nil)

func (s *SignUpService) SignUp(input dtos.SignupInput) error {
	_, err := s.repo.FindByEmail(input.Email)
	if err == nil {
		return errors.New("e-mail already used")
	}

	password, err := s.passwordGenerator.Generate()
	if err != nil {
		return err
	}

	hashedPassword, err := s.hasher.Hash(password)
	if err != nil {
		return err
	}

	user := models.User{
		Email:       input.Email,
		DateOfBirth: input.DateOfBirth,
		Name:        input.Name,
		Password:    hashedPassword,
	}

	subject := "Your First Access Password"
	body := fmt.Sprintf("Hello %s,\n\nYour first access password is: %s\n\nPlease login and change your password after your first access.", input.Name, password)
	err = s.emailSender.Send(user.Email, subject, body)
	if err != nil {
		return err
	}

	return nil
}
