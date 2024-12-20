package testutil

import (
	"crypto/rand"
)

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	_, err := rand.Read(token)
	if err != nil {
		panic(err)
	}

	return token
}
