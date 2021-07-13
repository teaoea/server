package tools

import (
	"crypto/sha256"
	"fmt"
)

func Checksum(body string) string {

	hash := sha256.New()
	hash.Write([]byte(body))
	value := fmt.Sprintf("%x", hash.Sum(nil))
	return value
}
