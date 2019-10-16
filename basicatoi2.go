package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		for i := 0; i <= 9; i += 1 {
			if rune(i) == r-'0' {
				result *= 10
				result += i
				break
			}
		}
	}
	return result
}
