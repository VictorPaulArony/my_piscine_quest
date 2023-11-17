package piscine

func IsNumeric(s string) bool {
	for _, num := range s {
		if !(num >= '0' && num <= '9') {
			return false
		}
	}
	return true
}
