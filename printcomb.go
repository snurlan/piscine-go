package piscine

import "github.com/01-edu/z01"

func PrintComb() {

	flag := false

	for i := '0'; i <= '9'; i += 1 {
		for j := i + 1; j <= '9'; j += 1 {
			for k := j + 1; k <= '9'; k += 1 {
				if flag {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				} else {
					flag = true
				}
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
			}
		}
	}

	z01.PrintRune('\n')
}
