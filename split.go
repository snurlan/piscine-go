package piscine

func Split(str, charset string) []string {
	n := 0
	m := 0
	for i := range str {
		n = i + 1
	}
	for i := range charset {
		m = i + 1
	}
	foundSep := true
	for foundSep {
		foundSep = false
		for i := 0; i < n-m; i++ {
			if str[i:i+m] == charset {
				foundSep = true
				str = str[:i] + " " + str[i+m:]
				n -= m
				break
			}
		}
	}

	return SplitWhiteSpaces(str)
}
