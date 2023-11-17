package piscine

func Join(strs []string, sep string) string {
	runes := ""
	count := 0
	for i := range strs {
		count = i + 1
	}
	for i, char := range strs {
		if i == count-1 {
			runes += char
		} else {
			runes += char + sep
		}
	}
	return runes
}
