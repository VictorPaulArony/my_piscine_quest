package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)

	item := ""
	for _, char := range str {
		if char != ' ' {
			item += string(char)
		} else {
			summary[item]++
			item = ""
		}
	}

	summary[item]++

	return summary
}

/*func main() {
	receipt := "apple apple banana orange apple banana"
	summary := countItems(receipt)

	for item, count := range summary {
		fmt.Printf("%s: %d\n", item, count)
	}
}*/
