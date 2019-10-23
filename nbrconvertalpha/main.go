package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	isUpper := false
	for i, arg := range args {
		if i == 0 && arg == "--upper" {
			isUpper = true
			z01.PrintRune(' ')
			continue
		}
		num := 0
		invalid := false
		for _, r := range arg {
			if r < '0' || r > '9' {
				invalid = true
				break
			}
			num = 10*num + int(r-'0')
		}
		if invalid || num == 0 || num > 26 {
			z01.PrintRune(' ')
			continue
		}
		var r rune
		if r = 'a' + rune(num-1); isUpper {
			r = 'A' + rune(num-1)
		}
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
