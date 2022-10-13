package goalHash

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

// Hash string to sha256
func Sha256(content string) string {
	hash := sha256.New()
	hash.Write([]byte(content))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Hash string to sha256 and encode with base64
func Sha256WithBase64(content string) string {
	hash := sha256.New()
	hash.Write([]byte(content))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// Hash string to sha512
func Sha512(content string) string {
	hash := sha512.New()
	hash.Write([]byte(content))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Hash string to sha256 and encode with base64
func Sha512WithBase64(content string) string {
	hash := sha512.New()
	hash.Write([]byte(content))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
