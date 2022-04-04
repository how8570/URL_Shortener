package v1

import (
	"math/rand"
)

const alphabate string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
const hashLen int = len(alphabate)

func Hash() string {
	// not real hash but just for demo
	var hash string
	for i := 0; i < 7; i++ {
		hash += string(alphabate[rand.Intn(hashLen)])
	}
	return hash
}
