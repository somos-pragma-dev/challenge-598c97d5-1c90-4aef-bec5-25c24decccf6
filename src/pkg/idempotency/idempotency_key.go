package idempotency

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateIdempotencyKey(req interface{}) string {
	// Implement idempotency key generation
	hash := sha256.New()
	hash.Write([]byte("some-data"))
	return hex.EncodeToString(hash.Sum(nil))
}

func IsIdempotent(key string) bool {
	// Implement idempotency check
	return false
}

func MarkAsIdempotent(key string) {
	// Implement idempotency mark
}