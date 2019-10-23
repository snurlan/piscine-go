package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	return r == 'a' || r == 'o' || r == 'u' || r == 'e' || r == 'i' || r == 'A' || r == 'O' || r == 'U' || r == 'E' || r == 'I'
}

func main() {
	args := os.Args[1:]
	var vowels []rune
	cnt := 0
	hasArgs := false
	for _, arg := range args {
		hasArgs = true
		for _, r := range arg {
			if isVowel(r) {
				vowels = append(vowels, r)
				cnt++
			}
		}
	}

	if !hasArgs {
		z01.PrintRune('\n')
		return
	}

	cur := 0
	isFirst := true
	for _, arg := range args {
		if isFirst {
			isFirst = false
		} else {
			z01.PrintRune(' ')
		}
		for _, r := range arg {
			if isVowel(r) {
				z01.PrintRune(vowels[cnt-cur-1])
				cur++
			} else {
				z01.PrintRune(r)
			}
		}
	}
}
