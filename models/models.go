package models

import "fmt"

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

func (c *Colony) Bfs() [][]int {
	toBeVisited := [][]int{}

	startPtr := c.Start
	endPtr := c.End
	startValue := *startPtr
	targetValue := *endPtr

	startRoom := startValue.RoomNumber
	endRoom := targetValue.RoomNumber
	var visitedRooms []int
	visitedRooms = append(visitedRooms, startRoom)

	for len(visitedRooms) != 0 {
		currentRoom := visitedRooms[len(visitedRooms)-1]
		
		if currentRoom == endRoom {
			fmt.Println(toBeVisited)
			
		}

		hasBeenVisited := make(map[int]bool)
		visitedVertex := visitedRooms[len(visitedRooms)-1]

		if _, exists := c.Rooms[visitedVertex]; exists {
			hasBeenVisited[visitedVertex] = true
			for _, connection := range c.Rooms[visitedVertex].Connection {
				if !hasBeenVisited[connection] {
					visitedRooms = append(visitedRooms, visitedVertex, connection)
					fmt.Println(visitedRooms)
					toBeVisited = append(toBeVisited, visitedRooms)
					toBeVisited = toBeVisited[1:]
				}
			}
		}
	}
	return toBeVisited
}

func (c *Colony) Dfs() []int {
	visited := []int{}
	toBeVisited := []int{}

	startPtr := c.Start
	endPtr := c.End
	startValue := *startPtr
	targetValue := *endPtr

	startRoom := startValue.RoomNumber
	endRoom := targetValue.RoomNumber

	toBeVisited = append(toBeVisited, startRoom)
	hasBeenVisited := make([]bool, len(c.Rooms))

	for len(toBeVisited) != 0 {
		if len(visited) > 0 && visited[len(visited)-1] == endRoom {
			break
		}
		visitedVertex := toBeVisited[0]
		for _, rooms := range c.Rooms {
			if rooms.RoomNumber == visitedVertex {
				hasBeenVisited[visitedVertex] = true
				for _, connection := range rooms.Connection {
					if !hasBeenVisited[connection] {
						visited = append(visited, connection)
						toBeVisited = append(toBeVisited, connection)
						toBeVisited = toBeVisited[1:]

					}
				}
			}
		}
	}

	return visited
}
