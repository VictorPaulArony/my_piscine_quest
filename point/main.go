package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func convertString(i int) {
	rn := '0'
	if (i / 10) > 0 {
		convertString(i / 10)
	}
	for j := 0; j < (i % 10); j++ {
		rn++
	}
	z01.PrintRune(rn)
}

func printVal(s string) {
	arr := []rune(s)
	for _, char := range arr {
		z01.PrintRune(char)
	}
}

func main() {
	points := &point{}
	setPoint(points)
	printVal("x = ")
	convertString(points.x)
	printVal(", y = ")
	convertString(points.y)
	z01.PrintRune('\n')
}
