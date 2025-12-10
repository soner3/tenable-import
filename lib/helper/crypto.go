package helper

import (
	"crypto/rand"
	"math/big"
)

// randomInputString is the set of characters used to generate random strings.
const randomInputString = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// GenerateRandomString returns a random string of length n.
func GenerateRandomString(n int) (string, error) {
	ret := make([]byte, n)
	for i := range n {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(randomInputString))))
		if err != nil {
			return "", err
		}
		ret[i] = randomInputString[num.Int64()]
	}

	return string(ret), nil
}
