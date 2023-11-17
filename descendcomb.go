package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	i := '9'
	j := '9'
	k := '9'
	l := '8'

	for i >= '0' {
		for j >= '0' {
			for k >= '0' {
				for l >= '0' {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(l)

					if !(i == '0' && j == '0' && k == '0' && l == '0') {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

					l--
				}
				k--
				l = '9'
			}
			j--
			k = '9'
			l = '9'
		}
		i--
		j = '9'
		k = '9'
		l = '9'
		if i < '0' {
			break
		}
	}
	z01.PrintRune('\n')
}

/*func main() {
	DescendComb()
}*/
