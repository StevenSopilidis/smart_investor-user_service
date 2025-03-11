package utils

import (
	"crypto/rand"
	"math/big"
)

type StringGenerator struct{}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"0123456789"

func (g *StringGenerator) Generate(length int) (string, error) {
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))

		if err != nil {
			return "", err
		}

		result[i] = charset[num.Uint64()]
	}

	return string(result), nil
}
