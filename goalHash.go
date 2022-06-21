package goalHash

import (
	"crypto/sha256"
	"encoding/base64"
)

// Hash string to sha256
func Sha256(content string) string {
	hashPasswordToSha256 := sha256.Sum256([]byte(content))
	passwordHashed := base64.StdEncoding.EncodeToString(hashPasswordToSha256[:])
	return passwordHashed
}

