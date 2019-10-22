package piscine

func ToLower(s string) string {
	runes := []rune(s)
	diff := 'A' - 'a'
	for i := range runes {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] -= diff
		}
	}
	return string(runes)
}
