package rand

import (
	"crypto/rand"
	"encoding/base64"
)

const RememberTokenBytes = 32

func RememberToken() (string, error) {
	return String(RememberTokenBytes)
}

//generate n random bytes, or will return an error if there was one.
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//base64 URL encoded version of that byte slice
func String(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}