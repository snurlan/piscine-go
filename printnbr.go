package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n/10 != 0 {
			PrintNbr(n / -10)
		}
		d := '0'
		for i := 0; i < -(n % 10); i += 1 {
			d += 1
		}
		z01.PrintRune(d)
	} else if n == 0 {
		z01.PrintRune('0')
	} else {
		if n/10 != 0 {
			PrintNbr(n / 10)
		}
		d := '0'
		for i := 0; i < n%10; i += 1 {
			d += 1
		}
		z01.PrintRune(d)
	}
}
