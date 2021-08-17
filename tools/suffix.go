package tools

import (
	"strings"

	"server/config/vars"
)

func SuffixCheck(email string) bool {
	addr := strings.Split(email, "@") // string segmentation
	suffix := "@" + addr[1]           // intercept email address suffix
	for _, value := range vars.EmailSuffixes {
		if value == suffix {
			return true
		}
	}
	return false
}
