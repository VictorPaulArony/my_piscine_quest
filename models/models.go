package models

import "fmt"

type Room struct {
	RoomNumber string
	Coordinate []int
	Connection []string
}
type Colony struct {
	NumAnts int
	Rooms   map[string]*Room
	Start   *Room
	End     *Room
}

func (c *Colony) Bfs() [][]string{
	toBeVisited := [][]string{}
	validPaths := [][]string{}

	startPtr := c.Start
	endPtr := c.End
	startValue := *startPtr
	targetValue := *endPtr

	startRoom := []string{startValue.RoomNumber}
	endRoom := targetValue.RoomNumber

	toBeVisited = append(toBeVisited, startRoom)
	currentRoom := startRoom

	for len(toBeVisited) != 0 {
		currentRoom = toBeVisited[0]
		toBeVisited = toBeVisited[1:]

		if currentRoom[len(currentRoom)-1] == endRoom {

			validPaths = append(validPaths, currentRoom)

			continue
		}
		hasBeenVisited := make(map[string]bool)

		for _, room := range currentRoom {
			hasBeenVisited[room] = true
		}
		if _, exists := c.Rooms[currentRoom[len(currentRoom)-1]]; exists {
			for _, connection := range c.Rooms[currentRoom[len(currentRoom)-1]].Connection {
				matrix := []string{}
				matrix = append(matrix, currentRoom...)
				if !hasBeenVisited[connection] {
					matrix = append(matrix, connection)
					toBeVisited = append(toBeVisited, matrix)

				}
			}
		}
	}
	return validPaths
}

func (c *Colony) Dfs() [][]string{
	toBeVisited := [][]string{}
	validPaths := [][]string{}

	startPtr := c.Start
	endPtr := c.End
	startValue := *startPtr
	targetValue := *endPtr

	startRoom := []string{startValue.RoomNumber}
	endRoom := targetValue.RoomNumber

	toBeVisited = append(toBeVisited, startRoom)
	currentRoom := startRoom

	for len(toBeVisited) != 0 {
		currentRoom = toBeVisited[len(toBeVisited)-1]
		toBeVisited = toBeVisited[:len(toBeVisited)-1]

		if currentRoom[len(currentRoom)-1] == endRoom {
			validPaths = append(validPaths, currentRoom)
			continue

		}
		hasBeenVisited := make(map[string]bool)
		for _, room := range currentRoom {
			hasBeenVisited[room] = true
		}
		if _, exists := c.Rooms[currentRoom[len(currentRoom)-1]]; exists {
			for _, connection := range c.Rooms[currentRoom[len(currentRoom)-1]].Connection {
				matrix := []string{}
				matrix = append(matrix, currentRoom...)
				if !hasBeenVisited[connection] {
					matrix = append(matrix, connection)
					toBeVisited = append(toBeVisited, matrix)
				}
			}
		}
	}
	fmt.Println("Dfs paths")
	return validPaths
}

// func (c *Colony) Dfs() []int {
// 	visited := []int{}
// 	toBeVisited := []int{}

// 	startPtr := c.Start
// 	endPtr := c.End
// 	startValue := *startPtr
// 	targetValue := *endPtr

// 	startRoom := startValue.RoomNumber
// 	endRoom := targetValue.RoomNumber

// 	toBeVisited = append(toBeVisited, startRoom)
// 	hasBeenVisited := make([]bool, len(c.Rooms))

// 	for len(toBeVisited) != 0 {
// 		if len(visited) > 0 && visited[len(visited)-1] == endRoom {
// 			break
// 		}
// 		currentRoom := toBeVisited[0]
// 		for _, rooms := range c.Rooms {
// 			if rooms.RoomNumber == currentRoom {
// 				hasBeenVisited[currentRoom] = true
// 				for _, connection := range rooms.Connection {
// 					if !hasBeenVisited[connection] {
// 						visited = append(visited, connection)
// 						toBeVisited = append(toBeVisited, connection)
// 						toBeVisited = toBeVisited[1:]

// 					}
// 				}
// 			}
// 		}
// 	}

// 	return visited
// }
