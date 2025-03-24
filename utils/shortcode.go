package utils

import (
    "crypto/rand"
    "math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const codeLength = 7

func GenerateShortCode() string {
	shortCode := make([]byte, codeLength)
	charSetLen := big.NewInt(int64(len(charset)))

	for i := range shortCode {
		randIndex, _ := rand.Int(rand.Reader, charSetLen)
		shortCode[i] = charset[randIndex.Int64()]
	}
	return string(shortCode)
}