package piscine

func StringToIntSlice(str string) []int {
	var slice []int

	for _, char := range str {
		num := int(char)
		slice = append(slice, num)
	}

	return slice
}

/*func main() {
	fmt.Println(StringToIntSlice("A quick brown fox jumps over the lazy dog"))
	fmt.Println(StringToIntSlice("Converted this string into an int"))
	fmt.Println(StringToIntSlice("hello THERE"))
}*/
