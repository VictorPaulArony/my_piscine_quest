package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	combinations "lemin/PathCombinations"
	"lemin/models"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Only one input-text file name needed")
		return
	}
	if !strings.HasSuffix(args[0], ".txt") {
		fmt.Println("the file given must be a .txt file")
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Error in opening file")
		return
	}
	defer file.Close()

	var colony models.Colony
	var startRoom, endRoom string
	var connectionFrom, connectionTo, roomPosition string
	var xCoordinate, yCoordinate,  antCount int

	colony.Rooms = make(map[string]*models.Room)
	scanner := bufio.NewScanner(file)

	var fileContent [][]string
	for scanner.Scan() {
		var room models.Room
		lineContent := scanner.Text()

		if len(lineContent) == 0 || string(lineContent[0]) == " " {
			continue
		}
		if strings.HasPrefix(lineContent, "#") {
			if strings.Contains(lineContent, "##start") {
				fileContent = append(fileContent, []string{"start"})
			} else if strings.Contains(lineContent, "##end") {
				fileContent = append(fileContent, []string{"end"})
			} else {
				continue
			}
		} else if strings.Contains(lineContent, " ") {
			roomDetails := strings.Split(lineContent, " ")
			fileContent = append(fileContent, roomDetails)
			roomPosition=roomDetails[0]
			xCoordinate,_=strconv.Atoi(roomDetails[1])
			yCoordinate,_=strconv.Atoi(roomDetails[2])

			room = models.Room{
				RoomNumber: roomPosition,
				Coordinate: []int{xCoordinate, yCoordinate},
			}
			colony.Rooms[roomPosition] = &room

		} else if strings.Contains(lineContent, "-") {

			connectionFromTo := strings.Split(lineContent, "-")

			connectionFrom = connectionFromTo[0]
			connectionTo = connectionFromTo[1]

			if _, exist1 := colony.Rooms[connectionFrom]; exist1 {
				colony.Rooms[connectionFrom].Connection = append(colony.Rooms[connectionFrom].Connection, connectionTo)
			}
			if _, exist2 := colony.Rooms[connectionTo]; exist2 {
				colony.Rooms[connectionTo].Connection = append(colony.Rooms[connectionTo].Connection, connectionFrom)
			}
		} else {
			antCount, _ = strconv.Atoi(lineContent)
		}

		for i := 0; i < len(fileContent)-1; i++ {
			if fileContent[i][0] == "end" {
				endRoom=fileContent[i+1][0]
				colony.End = colony.Rooms[endRoom]
				break
			}
		}
		colony.NumAnts = antCount
	}
	startRoom = fileContent[1][0]
	colony.Start = colony.Rooms[startRoom]


	bfsValidPaths := colony.Bfs()
	// dfsValidPaths:=colony.Dfs()
	antPaths := combinations.FindCombinations(bfsValidPaths)
	fmt.Println(antPaths)
}
