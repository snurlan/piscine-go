package main

import "github.com/01-edu/z01"

func main() {
	for x := 'z'; x >= 'a'; x -= 1 {
		z01.PrintRune(x)
	}
	z01.PrintRune('\n')
}
