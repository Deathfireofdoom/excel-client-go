package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateUUID() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	// Set version (4) and variant (2)
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// Format UUID as a string
	uuidString := hex.EncodeToString(uuid)
	return uuidString, nil
}
