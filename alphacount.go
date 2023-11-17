package piscine

func AlphaCount(s string) int {
	count := 0
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z' {
			count++
		}
	}
	return count
}
