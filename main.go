package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	count := len(args)
	msg := "odd"
	if count%2 == 0 {
		msg = "even"
	}
	printStr("I have an " + msg + " number of arguments")
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
