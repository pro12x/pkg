package pkg

import (
	"crypto/rand"
	"encoding/hex"
)

func Uppercase(s string) string {
	var result []byte
	if len(s) != 0 {
		for _, c := range s {
			if c >= 'a' && c <= 'z' {
				result = append(result, byte(c-32))
				continue
			}
			result = append(result, byte(c))
		}
	}
	return string(result)
}

func GenerateEncryptionKey() (string, error) {
	// Define the length of the key in bytes (32, 64, 128... => 256, 512, 1024...)
	size := 32

	key := make([]byte, size)

	// Fill the key slice with cryptographic keys
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	// Convert the key to hexadecimal
	result := hex.EncodeToString(key)

	return Uppercase(result), nil
}
