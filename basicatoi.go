package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, r := range s {
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
