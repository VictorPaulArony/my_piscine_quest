package piscine

func Any(f func(string) bool, a []string) bool {
	for _, elem := range a {
		if f(elem) {
			return true
		}
	}
	return false
}

/*func IsNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}*/
/*
)

func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := piscine.Any(piscine.IsNumeric, a1)
	result2 := piscine.Any(piscine.IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

And its output :

$ go run .
false
true
$
*/
