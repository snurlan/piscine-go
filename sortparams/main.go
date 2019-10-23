package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	ascii := [128]int{}
	runes := [128]rune{}
	for _, arg := range args {
		for _, r := range arg {
			ascii[r]++
			runes[r] = r
		}
	}
	for i := range ascii {
		for ascii[i] > 0 {
			z01.PrintRune(runes[i])
			z01.PrintRune('\n')
			ascii[i]--
		}
	}
}
