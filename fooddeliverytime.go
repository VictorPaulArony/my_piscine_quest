package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	menu := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}
	item, ok := menu[order]
	if !ok {
		return 404
	}
	return item.preptime
}

/*func main() {
	order1 := "burger"
	order2 := "chips"
	order3 := "salad"

	time1 := FoodDeliveryTime(order1)
	time2 := FoodDeliveryTime(order2)
	time3 := FoodDeliveryTime(order3)

	fmt.Println(order1, ":", time1, "minutes")
	fmt.Println(order2, ":", time2, "minutes")
	fmt.Println(order3, ":", time3, "minutes")
}*/
