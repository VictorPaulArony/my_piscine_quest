package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	RoomName string
	X, Y     int
	Links    []string
}

type Graph struct {
	TotalAnts int
	Rooms     map[string]*Room
	Start     Room
	End       Room
}

// function to read the .txt file to get the data
func ReadFile(fileName string) (*Graph, error) {
	// read the file to get the graph details
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error while reding the file", err)
		return nil, err
	}

	// close the file after the reading is complete
	defer file.Close()

	// initialize the graph to store the data
	graph := IntializeGraph()

	// scan through the file to populate the graph with the file data
	scanner := bufio.NewScanner(file)
	// check if the file is an empty file
	if !scanner.Scan() {
		log.Println("The file provided is empty")
		return nil, err
	}

	ants, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Println("Error while converting string to num for ants: ", err)
		return nil, err
	}

	// update the number of ants to the graph
	graph.TotalAnts = ants

	start := false
	end := false
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		// check for the start and end of the file rooms
		if line == "##start" {
			start = true
		} else if line == "##end" {
			end = true
		}

		if strings.HasPrefix(line, "#") {
			continue // Skip comments
		}

		// poputate the roooms and their cooordinates
		if strings.Contains(line, " ") {
			roomDetails := strings.Fields(line)

			roomName := roomDetails[0]

			// x and y coordinates for the rooms
			x, err1 := strconv.Atoi(roomDetails[1])
			y, err2 := strconv.Atoi(roomDetails[2])
			if err1 != nil || err2 != nil {
				log.Println("Error while converting string to num for x and y: ", err)
				return nil, err
			}

			room := &Room{
				RoomName: roomName,
				X:        x,
				Y:        y,
			}

			// populate the start and the end of the rooms to the graph
			if start {
				graph.Start = *room
				start = false
			} else if end {
				graph.End = *room
				end = false
			}

			// populate the room to the map of the rooms
			graph.Rooms[roomName] = room

			// going to the rooms and their links within the tunnel

		} else if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				log.Println("Error invalid room links format")
				return nil, err
			}
			// fmt.Printf("%q \n",parts[1])

			// Check if the rooms in the links are in the map of all the rooms
			room1, exist1 := graph.Rooms[parts[0]]
			room2, exist2 := graph.Rooms[parts[1]]

			if !exist1 || !exist2 {
				log.Println("The links provided does not exist")
				return nil, err
			}
			room1.Links = append(room1.Links, parts[1])
			room2.Links = append(room2.Links, parts[0])

		}

	}

	return graph, nil
}

// function to initialize the Graph
func IntializeGraph() *Graph {
	return &Graph{
		TotalAnts: 0,
		Rooms:     make(map[string]*Room),
	}
}
