package piscine

func LastRune(s string) rune {
	if len(s) > 0 {
		runes := []rune(s)
		lastIndex := len(runes) - 1
		return runes[lastIndex]
	}
	return 0
}
