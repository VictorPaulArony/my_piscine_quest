package piscine

import "github.com/01-edu/z01"

func IsNegative(num int) {
	if num < 0 {
		z01.PrintRune('T')
	}
	if num >= 0 {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
	return
}
