package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	diff := 'A' - 'a'
	for i := range runes {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] += diff
		}
	}
	return string(runes)
}
