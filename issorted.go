package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) > 1 {
		if f(a[0], a[1]) >= 0 {
			for i := 0; i < len(a)-1; i++ {
				if f(a[i], a[i+1]) < 0 {
					return false
				}
			}
		}
		if f(a[0], a[1]) <= 0 {
			for i := 0; i < len(a)-1; i++ {
				if f(a[i], a[i+1]) > 0 {
					return false
				}
			}
		}
	}

	return true
}

/*func main() {
	f := func(a, b int) int {
		return a - b
	}
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)

$ go run .
true
false
$


}*/
