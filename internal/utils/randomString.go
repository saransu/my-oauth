package utils

import (
	"math/rand"
)

const randomPool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewRandomString(length int) string {
	b := make([]byte, length)

	for i := range length {
		b[i] = randomPool[rand.Intn(len(randomPool))]
	}

	return string(b)
}
