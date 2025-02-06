package main

import (
	"fmt"
	"lem-in/utils"
)

func main() {

	// reading the file
	colony := utils.FileReader("test0.txt")

	fmt.Println("number_of_ants ", colony.NoOfAnts)
	rooms := colony.Rooms

	fmt.Print("the_rooms\n")
	var x, y int
	for _, room := range rooms{
		for _, v := range room.HouseAndCoordinates {
			x = v[0]
			y = v[1]
		}
		fmt.Println(room.Name, x, y)
	}
	
	fmt.Print("the_links\n")
	for _, room := range rooms {
		for k, v := range room.Link{
			fmt.Printf("%s-%s\n", k, v)
		}
	}
}
