package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"strings"
)

func CreateHash(key ...string) string {
	hasher := sha256.New()
	hasher.Write([]byte(strings.Join(key, "")))
	return hex.EncodeToString(hasher.Sum(nil))
}

//CreateSmallHash, generate hash and split hash result.
func CreateSmallHash(ln int, key ...string) string {
	res := CreateHash(key...)
	l := ln / 2
	return res[:l] + res[len(res)-l:]
}

//CreateRandomId, generate hash using uuid and split hash result.
func CreateRandomId(ln int) string {
	return CreateSmallHash(ln, uuid.New().String())
}
