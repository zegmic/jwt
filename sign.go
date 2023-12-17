package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
)

type HMACSHA256 struct {
	key []byte
}

func NewHMACSHA256Signer(key string) HMACSHA256 {
	return HMACSHA256{
		key: []byte(key),
	}
}

func (a HMACSHA256) Sign(payload string) []byte {
	h := hmac.New(sha256.New, a.key)
	h.Write([]byte(payload))

	return h.Sum(nil)
}

func (a HMACSHA256) Name() string {
	return "HS256"
}
