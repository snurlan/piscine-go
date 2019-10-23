package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for _, s := range table {
		for _, r := range []rune(s) {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
