package controllers

import (
	"msvc-users/domain/dtos"
	"msvc-users/domain/usecases"
	"msvc-users/presentation/contracts"
	validationContracts "msvc-users/validation/contracts"
)

type SignInController struct {
	signInUseCase usecases.SignIn
	validator     validationContracts.Validator
}

func NewSignInController(signInUseCase usecases.SignIn, validator validationContracts.Validator) *SignInController {
	return &SignInController{signInUseCase: signInUseCase, validator: validator}
}

var _ contracts.Controller = (*SignInController)(nil)

func (c *SignInController) Handle(httpRequest contracts.HttpRequest) (*contracts.HttpResponse, error) {
	var input dtos.SignInInput
	if httpRequest.Body != nil {
		input = httpRequest.Body.(dtos.SignInInput)
	}

	err := c.validator.Validate(input)
	if err != nil {
		return contracts.BadRequest(err), nil
	}

	token, err := c.signInUseCase.SignIn(input)
	if err != nil {
		return contracts.Unauthorized(err), nil
	}

	return contracts.Ok(token), nil
}
