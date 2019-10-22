package piscine

import "github.com/01-edu/z01"

func printNbrBase(nbr, n int, base []rune) {
	if nbr == 0 {
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
		printNbrBase(-(nbr / n), n, base)
		mod := nbr % n
		if mod < 0 {
			mod = -mod
		}
		z01.PrintRune(base[mod])
	} else {
		printNbrBase(nbr/n, n, base)
		z01.PrintRune(base[nbr%n])
	}
}

func PrintNbrBase(nbr int, base string) {
	rBase := []rune(base)
	n := 0
	isValid := true
	for i, r := range rBase {
		n = i + 1
		if r == '-' || r == '+' {
			isValid = false
			break
		}
	}

	for i := range rBase {
		for j := i + 1; j < n; j++ {
			if rBase[i] == rBase[j] {
				isValid = false
				break
			}
		}
	}

	if !isValid || n < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr == 0 {
		z01.PrintRune('0')
		return
	}

	printNbrBase(nbr, n, rBase)
}
