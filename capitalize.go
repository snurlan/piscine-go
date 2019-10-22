package piscine

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		r -= 'A' - 'a'
	}
	return r
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		r += 'A' - 'a'
	}
	return r
}

func isAlphanumeric(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9'
}

func Capitalize(s string) string {
	runes := []rune(s)
	flag := true
	for i, r := range runes {
		if !isAlphanumeric(r) {
			flag = true
			continue
		}
		if flag {
			runes[i] = toUpper(r)
			flag = false
		} else {
			runes[i] = toLower(r)
		}
	}
	return string(runes)
}
