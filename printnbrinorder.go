package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		z01.PrintRune('-')
		z01.PrintRune(rune(-(n % 10)) + '0')
		if n/-10 != 0 {
			PrintNbrInOrder(n / -10)
		}
	} else if n == 0 {
		z01.PrintRune('0')
	} else {
		z01.PrintRune(rune(n%10) + '0')
		if n/10 != 0 {
			PrintNbrInOrder(n / 10)
		}
	}
}
