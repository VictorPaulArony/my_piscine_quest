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
// using either BFS/DFS to get the best path and all paths
// sort the best path from the search
// deploy the ants to the paths
// show/print/simulate the ants in the paths
// end

func FileReader(path string) *models.Colony {

	var rooms []*models.Room
	var links map[string]string = make(map[string]string)
	var noOfAnts int
	

	file, err := os.Open(path)
	if err != nil {
		log.Println("Error opening the data file \n", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		// capture the number of ants
		ants := CheckNoOfAnts(text) //update this in the colony no of ants
		if ants > 0 {
			noOfAnts = ants
			continue
		}

		//Capture the rooms
		singleRoom := MapRooms(text)
		if singleRoom != nil {
			rooms = append(rooms, singleRoom) // update this on the colony as slice of rooms
		}

		newLinks := CheckForLinks(text)
		for k, v := range newLinks{
			links[k] = v
		}
	}

	//Assigning the links to their respective rooms
	for _, room := range rooms {
		if link, found := links[room.Name]; found {
			room.Link = map[string]string{room.Name: link}
		}
	}

	colony := &models.Colony{
		NoOfAnts: noOfAnts,
		Rooms: rooms,
	}

	return colony
}

// This function checks for single numbers and treats them as the number of ants.
func CheckNoOfAnts(text string) int {

	if strings.Contains(text, "-") || strings.HasPrefix(text, "##"){
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
	if strings.Contains(text, "-") || strings.HasPrefix(text, "##"){
		return nil
	}

	houseAndCoordinates := make(map[string][]int)
	var roomName string
	var coordinates []int

	splitted := strings.Split(text, " ")
	if len(splitted) == 3 {
		roomName = splitted[0]
	}
	x , err:= strconv.Atoi(splitted[1])
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
		Name: roomName,
		HouseAndCoordinates: houseAndCoordinates,
	}

	return room
}

// If the value of a particular line is separated by (-)
// we get them as a link to a house

func CheckForLinks(text string) map[string]string {
	roomLink := make(map[string]string)

	var splitted []string
	if strings.Contains(text, "-") {
		splitted = strings.Split(text, "-")
		if len(splitted) == 2 {
			roomLink[splitted[0]] = splitted[1]
		}
	}



	return roomLink
}
