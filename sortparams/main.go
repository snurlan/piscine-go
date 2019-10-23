package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	ascii := [128]int{}
	for _, arg := range args {
		for _, r := range arg {
			ascii[r]++
		}
	}
	for i := range ascii {
		for ascii[i] > 0 {
			z01.PrintRune(rune(i))
			z01.PrintRune('\n')
			ascii[i]--
		}
	}
}
