package tools

import "server/config/vars"

func SuffixCheck(email string) bool {
	for _, value := range vars.EmailSuffixes {
		if email == value {
			return true
		}
	}
	return false
}
