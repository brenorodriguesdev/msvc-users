package contracts

type Hasher interface {
	Hash(password string) (string, error)
}
