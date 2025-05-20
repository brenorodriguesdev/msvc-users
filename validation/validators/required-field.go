package validators

import (
	"errors"
	"fmt"
)

type RequiredFieldValidator struct {
	field string
}

func NewRequiredFieldValidator(field string) *RequiredFieldValidator {
	return &RequiredFieldValidator{field: field}
}

func (v *RequiredFieldValidator) Validate(data interface{}) error {
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return errors.New("invalid data format")
	}

	if _, exists := dataMap[v.field]; !exists {
		return fmt.Errorf("Missing param field: %s", v.field)
	}
	return nil
}
