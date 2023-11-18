package main

import (
	"os"

	"github.com/01-edu/z01"
)

func checkParam(args []string) bool {
	if len(args) != 9 {
		return false
	}
	for i := range args {
		if len(args[i]) != 9 {
			return false
		}
		for j := range args[i] {
			l := args[i][j]
			if l != '.' && (l < '1' || l > '9') {
				return false
			}
		}
	}
	return true
}

func printGreed(grid [][]byte) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if x > 0 {
				z01.PrintRune(' ')
			}
			z01.PrintRune(rune(grid[y][x]))
		}
		z01.PrintRune('\n')
	}
}

func checkSudoku(deep int, grid [][]byte) bool {
	x := deep % 9
	y := deep / 9
	// check line
	for i := 0; i < 9; i++ {
		if i != x && grid[y][i] == grid[y][x] {
			return false
		}
	}
	// check colone
	for i := 0; i < 9; i++ {
		if i != y && grid[i][x] == grid[y][x] {
			return false
		}
	}
	// check case
	caseX := x / 3
	caseY := y / 3
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			xx := caseX*3 + i
			yy := caseY*3 + j
			if (xx != x || yy != y) && grid[yy][xx] == grid[y][x] {
				return false
			}
		}
	}
	return true
}

func backTrack(deep int, grid [][]byte) bool {
	if deep == 81 {
		printGreed(grid)
		return true
	}
	x := deep % 9
	y := deep / 9
	if grid[y][x] != '.' {
		if checkSudoku(deep, grid) && backTrack(deep+1, grid) {
			return true
		}
	} else {
		for i := '1'; i <= '9'; i++ {
			grid[y][x] = byte(i)
			if checkSudoku(deep, grid) && backTrack(deep+1, grid) {
				return true
			}
		}
		grid[y][x] = '.'
	}
	return false
}

func resolveSudoku(args []string) bool {
	var grid [][]byte
	for i := range args {
		grid = append(grid, []byte(args[i]))
	}
	return backTrack(0, grid)
}

func main() {
	args := os.Args
	if !checkParam(args[1:]) || !resolveSudoku(args[1:]) {
		printLn("Error")
	}
	z01.PrintRune('\n')
}

func printLn(s string) {
	for _, runes := range s {
		z01.PrintRune(runes)
	}
}
