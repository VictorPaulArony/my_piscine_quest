package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"lemin/models"
)

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

	var room *models.Room
	//var colony *models.Colony
	var rooms []*models.Room
	var connectionFrom, connectionTo, roomPosition, xCoordinate, yCoordinate int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineContent := scanner.Text()

		if strings.Contains(lineContent, "##comment") {
			continue
		}
		if strings.Contains(lineContent, " ") {
			roomDetails := strings.Split(lineContent, " ")
			roomPosition, _ = strconv.Atoi(roomDetails[0])
			xCoordinate, _ = strconv.Atoi(roomDetails[1])
			yCoordinate, _ = strconv.Atoi(roomDetails[2])

			room = &models.Room{
				RoomNumber: roomPosition,
				Coordinate: []int{xCoordinate, yCoordinate},
			}

			rooms = append(rooms, room)
		}
		if strings.Contains(lineContent, "-") {

			connectionFromTo := strings.Split(lineContent, "-")

			connectionFrom, _ = strconv.Atoi(connectionFromTo[0])
			connectionTo, _ = strconv.Atoi(connectionFromTo[1])

			for _, room = range rooms {
				if room.RoomNumber == connectionFrom {
					room.Connection = append(room.Connection, connectionTo)
				} else {
					if room.RoomNumber == connectionTo {
						room.Connection = append(room.Connection, connectionFrom)
					}
				}
			}
		}
	}
	for _, room = range rooms {
		fmt.Println(room)
	}
}
