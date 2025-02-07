package main

import (
	"fmt"
	"log"
	"os"

	"lem-in/utils"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}

	fileName := os.Args[1]
	// g := utils.IntializeGraph()

	graph, err := utils.ReadFile(fileName)
	if err != nil {
		log.Println("Error in Reading file function")
		return
	}
	fmt.Printf("Total ants: %d\n", graph.TotalAnts)

	for _, rooms := range graph.Rooms {
		fmt.Printf("{ %s } links { %s } \n", rooms.RoomName, rooms.Links)
	}

	allpaths := graph.FindPaths(graph.Start.RoomName, graph.End.RoomName)
	fmt.Println(allpaths)
}
