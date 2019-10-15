package piscine

import "github.com/01-edu/z01"

func solve(n int, s []rune, m rune, isFirst *bool) {
	if n == 0 {
		if *isFirst {
			*isFirst = false
		} else {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		for i := 0; i < len(s); i += 1 {
			z01.PrintRune(s[i])
		}
	}

	for i := m; i <= '9'; i += 1 {
		solve(n-1, append(s, i), i+1, isFirst)
	}
}

func PrintCombN(n int) {
	isFirst := true
	solve(n, []rune{}, '0', &isFirst)
	z01.PrintRune('\n')
}
