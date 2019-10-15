package main

import "github.com/01-edu/z01"

func main() {
	for x := 'a'; x <= 'z'; x += 1 {
		z01.PrintRune(x)
	}
}
