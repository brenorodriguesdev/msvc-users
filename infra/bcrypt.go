package infra

import (
	"msvc-users/application/contracts"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct{}

func NewBcryptHasher() contracts.Hasher {
	return &BcryptHasher{}
}

var _ contracts.Hasher = (*BcryptHasher)(nil)

func (h *BcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

type BcryptHasherCompare struct{}

func NewBcryptHasherCompare() contracts.HasherCompare {
	return &BcryptHasherCompare{}
}

var _ contracts.HasherCompare = (*BcryptHasherCompare)(nil)

func (h *BcryptHasherCompare) Compare(password, hash string) (bool, error) {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil, nil
}
