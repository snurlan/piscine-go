package piscine

func NRune(s string, n int) rune {
	return []rune(s[n-1 : n])[0]
}
