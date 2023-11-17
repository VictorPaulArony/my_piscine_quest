package piscine

func Abort(a, b, c, d, e int) int {
	numbers := [5]int{a, b, c, d, e}

	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers[2]
}

/*func main() {
	a, b, c, d, e := 5, 3, 8, 2, 6
	median := Median(a, b, c, d, e)
	fmt.Println(median) // Output: 5
}*/
