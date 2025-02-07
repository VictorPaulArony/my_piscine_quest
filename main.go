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
	if !strings.HasSuffix(args[0], ".txt") {
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
	var startRoom, endRoom, antCount int
	var connectionFrom, connectionTo, roomPosition, xCoordinate, yCoordinate int
	var isEnd bool

	colony.Rooms = make(map[int]*Room)
	scanner := bufio.NewScanner(file)

	var fileContent [][]string
	for scanner.Scan() {
		lineContent := scanner.Text()

		if len(lineContent) == 0 || string(lineContent[0]) == " " {
			continue
		}
		if strings.HasPrefix(lineContent, "#") {
			if strings.Contains(lineContent, "##start") {
				fileContent = append(fileContent, []string{"start"})
			} else if strings.Contains(lineContent, "##end") {
				isEnd = true
			} else {
				continue
			}
		} else if strings.Contains(lineContent, " ") {
			roomDetails := strings.Split(lineContent, " ")
			fileContent = append(fileContent, roomDetails)
			roomPosition, _ = strconv.Atoi(roomDetails[0])
			xCoordinate, _ = strconv.Atoi(roomDetails[1])
			yCoordinate, _ = strconv.Atoi(roomDetails[2])

			room = &Room{
				RoomNumber: roomPosition,
				Coordinate: []int{xCoordinate, yCoordinate},
			}
			colony.Rooms[roomPosition] = room

		} else if strings.Contains(lineContent, "-") {

			connectionFromTo := strings.Split(lineContent, "-")

			connectionFrom, _ = strconv.Atoi(connectionFromTo[0])
			connectionTo, _ = strconv.Atoi(connectionFromTo[1])

			if _, exist1 := colony.Rooms[connectionFrom]; exist1 {
				colony.Rooms[connectionFrom].Connection = append(colony.Rooms[connectionFrom].Connection, connectionTo)
			}
			if _, exist2 := colony.Rooms[connectionTo]; exist2 {
				colony.Rooms[connectionTo].Connection = append(colony.Rooms[connectionTo].Connection, connectionFrom)
			}
		} else {
			antCount, _ = strconv.Atoi(lineContent)
		}

		_, exists1 := colony.Rooms[endRoom]
		if exists1 && isEnd {
			colony.End = room
			isEnd = false
		}

		colony.NumAnts = antCount
	}
	startRoom, _ = strconv.Atoi(fileContent[1][0])
	colony.Start = colony.Rooms[startRoom]
	fmt.Println(colony)
}
