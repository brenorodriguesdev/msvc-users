package factories

import (
	"msvc-users/validation/contracts"
	"msvc-users/validation/validators"
)

func MakeSignInValidator() contracts.Validator {
	requiredFields := []string{"email", "password"}
	var validations []contracts.Validator

	for _, field := range requiredFields {
		validations = append(validations, validators.NewRequiredFieldValidator(field))
	}

	return validators.NewValidatorComposite(validations)
}
