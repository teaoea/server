package tools

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	dig     = "1234567890"
	str     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	idxBits = 6
	idxMask = 1<<idxBits - 1
	idxMax  = 63 / idxBits
)

// RandomDig
/// generate a number of the specified length
func RandomDig(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = dig[rand.Int63()%int64(len(dig))]
	}
	return string(b)
}

// RandomStr
/// generate a string of the specified length
func RandomStr(length int) string {
	b := make([]byte, length)
	for i, cache, remain := length-1, rand.Int63(), idxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), idxMax
		}
		if idx := int(cache & idxMask); idx < len(str) {
			b[i] = str[idx]
			i--
		}
		cache >>= idxBits
		remain--
	}
	return string(b)
}
