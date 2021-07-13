package tools

import (
	"crypto/sha256"
	"fmt"
	"regexp"
)

func Checksum(body string) string {

	hash := sha256.New()
	hash.Write([]byte(body))
	value := fmt.Sprintf("%x", hash.Sum(nil))
	return value
}

func CheckPassword(password string) bool {
	num := `[0-9]{1}`
	a := `[a-z]{1}`
	A := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_<>\'\"]{1}`
	if (len(password)) < 8 || len(password) > 32 {
		return false
	}
	if ok, _ := regexp.MatchString(num, password); !ok {
		return false
	}

	if ok, _ := regexp.MatchString(a, password); !ok {
		return false
	}

	if ok, _ := regexp.MatchString(A, password); !ok {
		return false
	}

	if ok, _ := regexp.MatchString(symbol, password); !ok {
		return false
	}
	return true
}
