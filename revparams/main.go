// Write a program that prints the arguments received in the command line in reverse order.
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for arg := len(args) - 1; arg >= 0; arg-- {
		for _, char := range args[arg] {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
