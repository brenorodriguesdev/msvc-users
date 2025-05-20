package controllers

import (
	"msvc-users/domain/dtos"
	"msvc-users/domain/usecases"
	"msvc-users/presentation/contracts"
	validationContracts "msvc-users/validation/contracts"
)

type SignUpController struct {
	signUpUseCase usecases.SignUp
	validator     validationContracts.Validator
}

func NewSignUpController(signUpUseCase usecases.SignUp, validator validationContracts.Validator) *SignUpController {
	return &SignUpController{signUpUseCase: signUpUseCase, validator: validator}
}

var _ contracts.Controller = (*SignUpController)(nil)

func (c *SignUpController) Handle(httpRequest contracts.HttpRequest) (*contracts.HttpResponse, error) {
	var input dtos.SignupInput
	if httpRequest.Body != nil {
		input = httpRequest.Body.(dtos.SignupInput)
	}

	err := c.validator.Validate(input)
	if err != nil {
		return contracts.BadRequest(err), nil
	}

	err = c.signUpUseCase.SignUp(input)
	if err != nil {
		return contracts.BadRequest(err), nil
	}

	return contracts.Created(), nil
}
