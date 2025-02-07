package main

import (
	"fmt"
	"lem-in/utils"
)

func main() {

	// reading the file
	colony := utils.FileReader("test0.txt")

	fmt.Println("number_of_ants ", colony.NoOfAnts)

	fmt.Println("The starting room coordinate ", colony.StartRoom)
	fmt.Println("The ending room coordinate ", colony.EndRoom)

	// Anything partaining to the rooms
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

	// All the links that are available for the rooms.
	fmt.Print("the_links\n")
	for _, room := range colony.Link{
		fmt.Printf("%s-%s\n", room[0], room[1])
	}
	
	roomAndLinks := colony.RoomAndLinks

	fmt.Println(roomAndLinks)
}
