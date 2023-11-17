package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 5 {
		return 0
	} else if nb == 0 {
		return 1
	}
	vi := nb - 1
	for i := vi; i >= 1; i-- {
		nb *= i
	}
	return nb
}
