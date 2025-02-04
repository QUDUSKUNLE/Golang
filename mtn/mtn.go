package mtn

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func GeneratePIN(length int) (string, error) {
	const digits = "0123456789"
	var result strings.Builder

	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		result.WriteByte(digits[index.Int64()])
	}
	return result.String(), nil
}
