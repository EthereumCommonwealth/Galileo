package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"hash"
	"encoding/base64"
)

func NewHMAC(key string) HMAC {
	h := hmac.New(sha256.New, []byte(key))
	return HMAC{
		hmac: h,
	}
}

//Hash will hash the provided input string using HMAC with the secret key provided when the HMAC object was created
func (h HMAC) Hash(input string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(input))
	b := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)
}

type HMAC struct {
	hmac hash.Hash
}