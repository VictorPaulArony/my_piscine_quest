package piscine

func ToLower(s string) string {
	runes := ""
	for _, char := range s {
		if char >= 65 && char <= 90 {
			char += 32
		}
		runes += string(char)
	}
	return runes
}
