package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)

	for index, value := range runes {
		if index == n-1 {
			return value
		}
	}
	return 0
}
