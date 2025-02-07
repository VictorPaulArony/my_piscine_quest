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

func (c *Colony) Bfs() []int {
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

		// visitedVertex := toBeVisited[0]
		for _, rooms := range c.Rooms {
			if !hasBeenVisited[toBeVisited[0]] {
				visited = append(visited, toBeVisited[0])
			}

			if rooms.RoomNumber == toBeVisited[0] {
				hasBeenVisited[toBeVisited[0]] = true
				for _, connection := range rooms.Connection {
					if !hasBeenVisited[connection] {
						toBeVisited = append(toBeVisited, connection)
						fmt.Println("connection", connection, "visted rooms", hasBeenVisited, toBeVisited)
						toBeVisited = toBeVisited[1:]
					}
				}

			}
		}
		fmt.Println("visited", visited)
	}
	return visited
}
