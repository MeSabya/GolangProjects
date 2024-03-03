package toolkit

import (
	"crypto/rand"
	"encoding/base64"
)

const randomStringSource = "abcdefghijklmnopqrstuvxyz0123456789-+"

// Tools is the type used to instantiate this modue.
type Tools struct {
}

func (t *Tools) RandomString(n int) (string, error) {

	bytesNeeded := (n * 6) / 8
	if (n*6)%8 != 0 {
		bytesNeeded++
	}

	randomBytes := make([]byte, bytesNeeded)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomString := base64.URLEncoding.EncodeToString(randomBytes)

	return randomString[:n], nil

}
