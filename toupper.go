package piscine

func ToUpper(s string) string {
	runes := ""
	for _, char := range s {
		if char >= 97 && char <= 122 {
			char -= 32
		}
		runes += string(char)
	}
	return runes
}
