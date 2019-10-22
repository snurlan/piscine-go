package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	cur := 0
	for index := range s {
		if index+1 == n {
			return runes[cur]
		}
		cur++
	}
	return '('
}
