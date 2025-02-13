package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"lem-in/models"
)

// reading the .txt file

func FileReader(fileNmae string) *models.Colony {
	var rooms []*models.Room                      // Holds all the rooms and their properties
	var allLinks [][]string                       // This holds all the links that are present in the data
	var roomAndConnectedLinks map[string][]string // Holds room and their connections
	// var noOfAnts int

	// Starting and ending coordinates
	var start *models.Room
	var end *models.Room

	// error checkers in the file reading and handling the error output
	countStart := 0
	countEnd := 0
	hasStartCoords := false
	hasEndCoords := false
	// Track seen rooms and links to prevent duplicates
	seenRooms := make(map[string]bool)
	seenLinks := make(map[string]bool)

	file, err := os.Open(fileNmae)
	if err != nil {
		fmt.Println("ERROR: invalid data format, file does not exist")
		return nil
	}
	defer file.Close()

	// var isStart, isEnd bool // Checks the first and end coordinates respectively
	scanner := bufio.NewScanner(file)

	// Read number of ants first
	if !scanner.Scan() {
		fmt.Println("ERROR: invalid data format, empty file")
		return nil
	}

	noOfAnts, err := strconv.Atoi(scanner.Text())
	if err != nil || noOfAnts < 1 {
		fmt.Println("ERROR: invalid data format, invalid number of ants")
		return nil
	}

	isStart := false
	isEnd := false

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			continue
		}

		// Handle commands
		if text == "##start" {
			countStart++
			if countStart > 1 {
				fmt.Println("ERROR: invalid data format, multiple start rooms")
				return nil
			}
			isStart = true
			continue
		}

		if text == "##end" {
			countEnd++
			if countEnd > 1 {
				fmt.Println("ERROR: invalid data format, multiple end rooms")
				return nil
			}
			isEnd = true
			continue
		}

		// Ignore any other instance of another #
		if strings.Contains(text, "#") {
			continue
		}

		// This holds all the rooms and their connected links
		if strings.Contains(text, "-") {
			links := strings.Split(text, "-")
			if len(links) != 2 {
				fmt.Println("ERROR: invalid data format, invalid link format")
				return nil
			}

			// Validate self-links
			if links[0] == links[1] {
				fmt.Println("ERROR: invalid data format, cannot link room to itself")
				return nil
			}

			// Check for duplicate links
			linkKey := fmt.Sprintf("%s-%s", links[0], links[1])
			reverseLinkKey := fmt.Sprintf("%s-%s", links[1], links[0])
			if seenLinks[linkKey] || seenLinks[reverseLinkKey] {
				fmt.Println("ERROR: invalid data format, duplicate link")
				return nil
			}
			seenLinks[linkKey] = true

			// Verify rooms exist
			if !roomExists(links[0], rooms) || !roomExists(links[1], rooms) {
				fmt.Println("ERROR: invalid data format, invalid room name (doesn't exist)")
				return nil
			}

			allLinks = append(allLinks, links)
			continue
		} else if strings.Contains(text, " ") {
			// Handle rooms
			roomData := strings.Split(text, " ")

			// Validate room name
			if strings.HasPrefix(roomData[0], "#") || strings.HasPrefix(roomData[0], "L") {
				fmt.Println("ERROR: invalid data format, invalid room name (L or #)")
				return nil
			}

			// Check for duplicate rooms
			if seenRooms[roomData[0]] {
				fmt.Println("ERROR: invalid data format, duplicate room")
				return nil
			}
			seenRooms[roomData[0]] = true

			room := MapRooms(text)
			if room == nil {
				fmt.Println("ERROR: invalid data format, invalid room coordinates")
				return nil
			}

			if isStart {
				start = room
				hasStartCoords = true
				isStart = false
			} else if isEnd {
				end = room
				hasEndCoords = true
				isEnd = false
			}

			rooms = append(rooms, room)
		}
	}

	// Final validations
	if !hasStartCoords {
		fmt.Println("ERROR: invalid data format, missing start room coordinates")
		return nil
	}
	if !hasEndCoords {
		fmt.Println("ERROR: invalid data format, missing end room coordinates")
		return nil
	}
	if countStart != 1 {
		fmt.Println("ERROR: invalid data format or no start room")
		return nil
	}
	if countEnd != 1 {
		fmt.Println("ERROR: invalid data format or no end room")
		return nil
	}
	if len(seenLinks) == 0 {
		fmt.Println("ERROR: invalid data format or no links found")
		return nil
	}

	// Create room connections map
	roomAndConnectedLinks = FindRoomAndLinks(allLinks)

	// Create and populate colony
	colony := &models.Colony{
		NoOfAnts:  noOfAnts,
		Rooms:     rooms,
		StartRoom: *start,
		EndRoom:   *end,
	}

	// Populate room links
	for _, room := range colony.Rooms {
		room.Link = allLinks
		room.RoomAndConnectedLinks = roomAndConnectedLinks
	}

	return colony
}

// This function checks for single numbers and treats them as the number of ants.
// func CheckNoOfAnts(text string) (int, error) {
// 	ants, err := strconv.Atoi(text)
// 	if err != nil || ants < 1 {
// 		log.Println("ERROR: invalid data format, error coneverting No of ants.\n", err)
// 		return 0, err
// 	}
// 	return ants, nil
// }

// if the length of a text is 3 after splitting with a space
// we get that as a room and the coordinates
func MapRooms(text string) *models.Room {
	if strings.Contains(text, "-") || strings.HasPrefix(text, "##") {
		return nil
	}
	houseAndCoordinates := make(map[string][]int)
	var roomName string
	var coordinates []int

	splitted := strings.Split(text, " ")

	if len(splitted) == 3 {
		roomName = splitted[0]
	}
	x, err := strconv.Atoi(splitted[1])
	if err != nil {
		log.Println("Error converting x coordiante")
		return nil
	}
	y, err := strconv.Atoi(splitted[2])
	if err != nil {
		log.Println("Error converting y coordiante")
		return nil
	}
	coordinates = append(coordinates, x, y)

	houseAndCoordinates[roomName] = coordinates

	room := &models.Room{
		Name:                roomName,
		HouseAndCoordinates: houseAndCoordinates,
	}

	return room
}

// // If the value of a particular line is separated by (-)
// // we get them as a link to a house

// func CheckForLinks(text string) []string {
// 	var roomLinks []string
// 	splitted := strings.Split(text, "-")
// 	if len(splitted) == 2 {
// 		roomLinks = (splitted)
// 	} else {
// 		log.Println("Invalid input")
// 		return nil
// 	}

// 	return roomLinks
// }

// // The rooms and the links that are connected to the room
func FindRoomAndLinks(c [][]string) map[string][]string {
	roomAndLinks := make(map[string][]string)
	for _, room := range c {
		roomAndLinks[room[0]] = append(roomAndLinks[room[0]], room[1])
		roomAndLinks[room[1]] = append(roomAndLinks[room[1]], room[0])
	}

	return roomAndLinks
}

// Helper function to check if a room exists
func roomExists(name string, rooms []*models.Room) bool {
	for _, room := range rooms {
		if room.Name == name {
			return true
		}
	}
	return false
}
