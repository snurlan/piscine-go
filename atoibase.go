package piscine

func AtoiBase(s string, base string) int {
	res := 0
	rBase := []rune(base)
	n := 0
	for i, r := range rBase {
		n = i + 1
		if r == '-' || r == '+' {
			return 0
		}
	}

	for i := range rBase {
		for j := i + 1; j < n; j++ {
			if rBase[i] == rBase[j] {
				return 0
			}
		}
	}

	if n < 2 {
		return 0
	}

	for _, r := range []rune(s) {
		for i, b := range rBase {
			if r == b {
				res = n*res + i
				break
			}
		}
	}

	return res
}
