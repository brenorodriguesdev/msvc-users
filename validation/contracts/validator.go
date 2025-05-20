package contracts

type Validator interface {
	Validate(data interface{}) error
}
