package main

import (
	"fmt"

	"lem-in/utils"
)

func areSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func findKeyForSlice(myMap map[string][]int, targetSlice []int) (string, bool) {
	// fmt.Printf("%v compared aganist %v\n", myMap, targetSlice)
	for key, slice := range myMap {
		// fmt.Printf("Key %s Value %v\n", key, slice)
		if areSlicesEqual(slice, targetSlice) {
			return key, true
		}
	}
	return "", false
}

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
	for _, room := range rooms {
		for _, v := range room.HouseAndCoordinates {
			x = v[0]
			y = v[1]
		}
		fmt.Println(room.Name, x, y)
	}

	// All the links that are available for the rooms.
	fmt.Print("the_links\n")
	for _, room := range colony.Link {
		fmt.Printf("%s-%s\n", room[0], room[1])
	}

	roomAndLinks := colony.RoomAndLinks

	fmt.Println(roomAndLinks)

	// The BFS shortest path

	my := make(map[string][]int)
	var startkey string
	var endkey string
	roomsAndCoord := colony.Rooms
	startSlice := colony.StartRoom
	endSlice := colony.EndRoom

	// Finding the starting point for the BFS
	for _, roomAnd := range roomsAndCoord {
		my = roomAnd.HouseAndCoordinates

		foundstartkey, found := findKeyForSlice(my, startSlice)
		if found {
			fmt.Println("START KEY ", foundstartkey)
			startkey = foundstartkey
			break
		} else {
			// fmt.Println("no start key")
			continue
		}
	}

	// Finding the ending place for the BFS
	for _, roomAnd := range roomsAndCoord{
		my = roomAnd.HouseAndCoordinates
		foundendkey, found := findKeyForSlice(my, endSlice)
		if found {
			fmt.Println("END KEY ", foundendkey)
			endkey = foundendkey
			break
		} else {
			// fmt.Println("No end key")
			continue
		}
	}

	shortestPath := utils.BfsShortestPath(startkey, endkey, roomAndLinks)

	if shortestPath != nil {
		fmt.Printf("Shortest path from %s to %s: %v\n", startkey, endkey, shortestPath)
	} else {
		fmt.Printf("No path found from %s to %s\n", startkey, endkey)
	}
}
