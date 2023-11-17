package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	slice := make([]int, max-min)
	for i := range slice {
		slice[i] = min + i
	}
	return slice
}
