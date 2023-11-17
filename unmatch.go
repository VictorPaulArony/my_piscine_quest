package piscine

func Unmatch(a []int) int {
	for _, nums := range a {
		counts := 0
		for _, num := range a {
			if num == nums {
				counts++
			}
		}
		if counts == 1 || counts%2 == 1 {
			return nums
		}
	}
	return -1
}
