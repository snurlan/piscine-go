package piscine

func IsUpper(str string) bool {
	runes := []rune(str)
	for _, r := range runes {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}
