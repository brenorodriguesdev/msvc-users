package utils

import (
	"crypto/rand"
	"encoding/base64"
	"msvc-users/application/contracts"
)

type RandomPasswordGenerator struct{}

func NewRandomPasswordGenerator() contracts.PasswordGenerator {
	return &RandomPasswordGenerator{}
}

var _ contracts.PasswordGenerator = (*RandomPasswordGenerator)(nil)

func (g *RandomPasswordGenerator) Generate() (string, error) {
	b := make([]byte, 12)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
