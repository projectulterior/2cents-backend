package utils

import (
	"crypto/rand"
	"fmt"
)

const alphanum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
const digits = "0123456789"

// Returns a random array of alpha-num of specified size
//
func Random(size uint) string {
	return random(size, alphanum)
}

func RandomDigits(size uint) string {
	return random(size, digits)
}

// Returns a random array of alpha-num of specified size
//
func random(size uint, characters string) string {
	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		panic(fmt.Sprintf("firestore: crypto/rand.Read error: %v", err))
	}
	for i, byt := range b {
		b[i] = characters[int(byt)%len(characters)]
	}
	return string(b)
}
