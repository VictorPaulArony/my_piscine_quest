package main

import "github.com/01-edu/z01"

func main() {
	for r := 122; r >= 97; r-- {
		z01.PrintRune(rune(r))
	}
	z01.PrintRune('\n')
}
