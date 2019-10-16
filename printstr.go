package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	for i := 0; i < len(str); i += 1 {
		z01.PrintRune(rune(str[i]))
	}
}
