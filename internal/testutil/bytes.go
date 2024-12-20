package testutil

import (
	"crypto/rand"
)

// RandomBytes returns random byte slice of given size.
func RandomBytes(size int) []byte {
	token := make([]byte, size)
	_, _ = rand.Read(token)

	return token
}
