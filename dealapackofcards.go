package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		fmt.Printf("Player %d: ", i+1)
		for j := i * 3; j < (i+1)*3; j++ {
			fmt.Printf("%d", deck[j])
			if j != (i+1)*3-1 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
