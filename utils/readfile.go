package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"lem-in/models"
)

// reading the .txt file

func FileReader(path string) *models.Colony {
	var rooms []*models.Room // Holds all the rooms and their properties
	var allLinks [][]string // This holds all the links that are present in the data
	var roomAndConnectedLinks map[string][]string //Holds room and their connections
	var noOfAnts int 

	// Starting and ending coordinates
	// var startCoord []int
	var start string
	// var endCoord []int
	var end string

	file, err := os.Open(path)
	if err != nil {
		log.Println("Error opening the data file \n", err)
		return nil
	}
	defer file.Close()

	var isStart, isEnd bool // Checks the first and end coordinates respectively
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if text == "##start" {
			isStart = true
			continue
		}
		if text == "##end" {
			isEnd = true
			continue
		}

		// Ignore any other instance of another #
		if strings.Contains(text, "#") {
			continue
		}

		// capture the number of ants
		ants := CheckNoOfAnts(text) // update this in the colony no of ants
		if ants > 0 {
			noOfAnts = ants
			continue
		}

		if isStart {
			startRoom := MapRooms(text)
			isStart = false
			// Get the room which will be the start
			for k := range startRoom.HouseAndCoordinates {
				start = k
			}
		} else if isEnd {
			endRoom := MapRooms(text)
			isEnd = false
			// Get the room which will be the end
			for k:= range endRoom.HouseAndCoordinates {
				end = k
			}
		}
		// Capture the rooms
		singleRoom := MapRooms(text)
		if singleRoom != nil {
			rooms = append(rooms, singleRoom) // update this on the colony as slice of rooms
		}
		// Capture links
		if strings.Contains(text, "-") {
			links := CheckForLinks(text)
			allLinks = append(allLinks, links)
		}
	}

	roomAndConnectedLinks = FindRoomAndLinks(allLinks) // This holds all the rooms and their connected links

	// Updating the colony struct
	colony := &models.Colony{
		NoOfAnts:     noOfAnts,
		Rooms:        rooms,
		StartRoom:    start,
		EndRoom:      end,
	}
	// Populate the rooms struct
	for _, room := range colony.Rooms{
		room.Link = allLinks
		room.RoomAndConnectedLinks = roomAndConnectedLinks
	}
	return colony
}

// This function checks for single numbers and treats them as the number of ants.
func CheckNoOfAnts(text string) int {
	if strings.Contains(text, "-") || strings.HasPrefix(text, "##") {
		return 0
	}

	splitted := strings.Split(text, " ")
	if len(splitted) == 1 {
		ants, err := strconv.Atoi(splitted[0])
		if err != nil {
			log.Println("Error coneverting No of ants.\n", err)
			return 0
		}
		return ants
	}

	return 0
}

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

// If the value of a particular line is separated by (-)
// we get them as a link to a house

func CheckForLinks(text string) []string {
	var roomLinks []string

	splitted := strings.Split(text, "-")
	if len(splitted) == 2 {
		roomLinks = (splitted)
	} else {
		log.Println("Invalid input")
		return nil
	}

	return roomLinks
}

// The rooms and the links that are connected to the room
func FindRoomAndLinks(c [][]string) map[string][]string {
	roomAndLinks := make(map[string][]string)
	for _, room := range c {
		roomAndLinks[room[0]] = append(roomAndLinks[room[0]], room[1])
		roomAndLinks[room[1]] = append(roomAndLinks[room[1]], room[0])
	}

	return roomAndLinks
}
