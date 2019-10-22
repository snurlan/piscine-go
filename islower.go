package piscine

func IsLower(str string) bool {
	runes := []rune(str)
	for _, r := range runes {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}
