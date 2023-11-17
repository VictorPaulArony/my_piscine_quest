package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, elem := range tab {
		if f(elem) {
			count++
		}
	}
	return count
}

/*func IsNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}*/
/*
func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This","1", "is", "4", "you"}
	answer1 := piscine.CountIf(piscine.IsNumeric, tab1)
	answer2 := piscine.CountIf(piscine.IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}

And its output :

$ go run .
0
2
$

*/
