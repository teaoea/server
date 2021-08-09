package tools

func SuffixCheck(email string) bool {
	for _, value := range []string{"", ""} {
		if email == value {
			return true
		}
	}
	return false
}
