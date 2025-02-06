package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	RoomNumber int
	Coordinate []int
	Connection []int
}
type Colony struct {
	NumAnts int
	Rooms   map[int]*Room
	Start   *Room
	End     *Room
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Only one input-text file name needed")
		return
	}
	if !strings.HasSuffix(args[0], "txt") {
		fmt.Println("the file given must be a .txt file")
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Error in opening file")
		return
	}

	var room *Room
	var colony Colony
	// var startRoom, endRoom int
	// var startPtr, endPtr *Room
	var connectionFrom, connectionTo, roomPosition, xCoordinate, yCoordinate int
	var fullFile [][]string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineContent := scanner.Text()

		if strings.Contains(lineContent, "#") {
			continue
		}
		if strings.Contains(lineContent, "##start") {
			startArr := []string{"##start"}
			fullFile = append(fullFile, startArr)
		}
		if strings.Contains(lineContent, "##end") {
			endArr := []string{"##end"}
			fullFile = append(fullFile, endArr)
		}
		if strings.Contains(lineContent, " ") {
			roomDetails := strings.Split(lineContent, " ")
			fullFile = append(fullFile, roomDetails)
			roomPosition, _ = strconv.Atoi(roomDetails[0])
			xCoordinate, _ = strconv.Atoi(roomDetails[1])
			yCoordinate, _ = strconv.Atoi(roomDetails[2])

			room = &Room{
				RoomNumber: roomPosition,
				Coordinate: []int{xCoordinate, yCoordinate},
			}

			colony = Colony{
				Rooms: map[int]*Room{},
			}

			colony.Rooms[roomPosition]= room
			
			for _, room := range colony.Rooms {
				fmt.Println(room)
			}
		}
		
		if strings.Contains(lineContent, "-") {
			
			connectionFromTo := strings.Split(lineContent, "-")
			
			connectionFrom, _ = strconv.Atoi(connectionFromTo[0])
			connectionTo, _ = strconv.Atoi(connectionFromTo[1])
			
			if _, exist1 := colony.Rooms[connectionFrom];exist1{
				colony.Rooms[connectionFrom].Connection= append(colony.Rooms[connectionFrom].Connection, connectionTo)
			}
			if _, exist2 := colony.Rooms[connectionTo];exist2 {
				colony.Rooms[connectionTo].Connection= append(colony.Rooms[connectionTo].Connection, connectionFrom)
			}
		}
		

		// antCount, _ := strconv.Atoi(lineContent)
		// for i := 0; i < len(fullFile)-1; i++ {
		// 	if fullFile[i][0] == "##start" {
		// 		startRoom, _ = strconv.Atoi(fullFile[i+1][0])
		// 	}
		// 	if fullFile[i][0] == "##end" {
		// 		endRoom, _ = strconv.Atoi(fullFile[i+1][0])
		// 	}
		// }

		// if _, exists := colony.Rooms[startRoom]; exists {
		// 	startPtr = room
		// }
		// if _, exists := colony.Rooms[endRoom]; exists {
		// 	endPtr = room
		// }

		// colony.NumAnts = antCount
		// colony.Start = startPtr
		// colony.End = endPtr
		
	}
	}
