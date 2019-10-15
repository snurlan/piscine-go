package piscine

import "github.com/01-edu/z01"

func PrintComb2() {

	flag := false

	for i := '0'; i <= '9'; i += 1 {
		for j := '0'; j <= '9'; j += 1 {
			for k := i; k <= '9'; k += 1 {
				for l := j; l <= '9'; l += 1 {
					if i == k && j == l {
						continue
					}
					if flag {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else {
						flag = true
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(l)
				}
			}
		}
	}

	z01.PrintRune('\n')
}
