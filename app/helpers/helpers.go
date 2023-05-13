package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func GenerateSHA256Hash(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	hash := h.Sum(nil)
	output := fmt.Sprintf("%x", string(hash[:]))

	return output
}
