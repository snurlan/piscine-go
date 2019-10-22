package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n/-10 != 0 {
			PrintNbrInOrder(n / -10)
		}
		z01.PrintRune(rune(-(n % 10)) + '0')
	} else if n == 0 {
		z01.PrintRune('0')
	} else {
		if n/10 != 0 {
			PrintNbrInOrder(n / 10)
		}
		z01.PrintRune(rune(n%10) + '0')
	}
}
