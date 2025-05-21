package contracts

type Encrypter interface {
	Encrypt(values interface{}) (string, error)
}
