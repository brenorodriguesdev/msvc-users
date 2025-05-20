package contracts

type PasswordGenerator interface {
	Generate() (string, error)
}
