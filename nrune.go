package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	return runes[(n-1)%len(runes)]
}
