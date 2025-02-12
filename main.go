package main

import (
	"fmt"
	"os"

	"lem-in/utils"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run . [FILE]")
		return
	}
	file := args[1]

	colony := utils.FileReader(file)
	fmt.Println("no_ants ", colony.NoOfAnts)
	fmt.Println("Start Room ", colony.StartRoom)
	fmt.Println("End Room ", colony.EndRoom)

	the_rooms := make(map[string][]int)
	// var the_links [][]string
	for _, room := range colony.Rooms {
		for r, coord := range room.HouseAndCoordinates {
			the_rooms[r] = coord
		}
	}

	fmt.Println("the_rooms")
	for r, coord := range the_rooms {
		fmt.Printf("%v %v %v\n", r, coord[0], coord[1])
	}
	fmt.Println("the_links")
	for _, room := range colony.Rooms {
		for _, v := range room.Link {
			fmt.Printf("%v-%v\n", v[0], v[1])
		}
		break
	}

	var roomAndConnections = map[string][]string{}
	for _, room := range colony.Rooms{
		for k, v := range room.RoomAndConnectedLinks{
			roomAndConnections[k] = v
		}
	}
	for k, v := range roomAndConnections{
		fmt.Printf("%s connected to %v\n", k, v)
	}

	colony.Bfs()
}
