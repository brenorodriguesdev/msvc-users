package validators

import (
	"msvc-users/validation/contracts"
)

type ValidatorComposite struct {
	validations []contracts.Validator
}

func NewValidatorComposite(validations []contracts.Validator) *ValidatorComposite {
	return &ValidatorComposite{validations: validations}
}

func (v *ValidatorComposite) Validate(data interface{}) error {
	for _, validator := range v.validations {
		err := validator.Validate(data)
		if err != nil {
			return err
		}
	}
	return nil
}
