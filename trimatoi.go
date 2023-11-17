package piscine

func TrimAtoi(s string) int {
	k := 1
	num := 0

	for _, a := range s {
		if a >= '0' && a <= '9' {
			b := int(a - 48)
			num = num*10 + b
		} else if a == '-' && num == 0 {
			k = -1
		}
	}
	return num * k
}
