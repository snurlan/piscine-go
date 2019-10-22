package piscine

func isAlpha(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func AlphaCount(str string) int {
	runes := []rune(str)
	cnt := 0
	for _, r := range runes {
		if isAlpha(r) {
			cnt++
		}
	}
	return cnt
}
