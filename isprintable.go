package piscine

func IsPrintable(str string) bool {
	runes := []rune(str)
	for _, r := range runes {
		if r == '\t' || r == '\n' || r == '\v' || r == '\f' || r == '\r' || r == ' ' {
			return false
		}
	}
	return true
}
