package main

import "fmt"

func main() {
	for x := 'a'; x <= 'z'; x += 1 {
		fmt.Print(string(x))
	}
}