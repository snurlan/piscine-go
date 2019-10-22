package piscine

func Index(s string, toFind string) int {
	sRunes := []rune(s)
	toFindRunes := []rune(toFind)
	for i := range sRunes {
		matches := true
		for j, r2 := range toFindRunes {
			if i+j >= len(sRunes) || sRunes[i+j] != r2 {
				matches = false
				break
			}
		}
		if matches {
			return i
		}
	}
	return -1
}
