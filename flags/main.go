package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("    This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("    This flag will behave like a boolean, if it is called it will order the argument.")
}

func printSorted(s string) {
	cnt := [1024]int{}
	for _, r := range s {
		cnt[r]++
	}
	for i := range cnt {
		for cnt[i] > 0 {
			z01.PrintRune(rune(i))
			cnt[i]--
		}
	}
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:]
	hasArgs := false
	needOrder := false
	var s = ""
	var inserts []string
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			printHelp()
			return
		}
		if arg == "-o" || arg == "--order" {
			needOrder = true
			continue
		}

		arg2 := arg + "xxxxxxxx"
		if arg2[:2] == "-i" {
			inserts = append(inserts, arg[3:])
		}
		if arg2[:8] == "--insert" {
			inserts = append(inserts, arg[9:])
		}

		hasArgs = true
		s = arg
	}
	if !hasArgs {
		printHelp()
		return
	}
	for _, insert := range inserts {
		s += insert
	}
	if needOrder {
		printSorted(s)
	} else {
		fmt.Println(s)
	}
}
