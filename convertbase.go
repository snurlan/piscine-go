package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	b1, b2 := 0, 0
	r1, r2 := []rune(baseFrom), []rune(baseTo)
	for i := range r1 {
		b1 = i + 1
	}
	for i := range r2 {
		b2 = i + 1
	}
	n := 0
	r0 := []rune(nbr)
	for _, d := range r0 {
		for i, r := range r1 {
			if d == r {
				n = n*b1 + i
			}
		}
	}
	if n == 0 {
		return string(r2[0])
	}
	res := ""
	for n > 0 {
		res = string(r2[n%b2]) + res
		n = n / b2
	}
	return res
}
