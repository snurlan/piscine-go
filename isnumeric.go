package piscine

func IsNumeric(str string) bool {
	runes := []rune(str)
	for _, r := range runes {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
