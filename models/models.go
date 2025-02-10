package models

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

		visitedVertex := toBeVisited[0]
		for _, rooms := range c.Rooms {
			if rooms.RoomNumber == visitedVertex {
				visited = append(visited, visitedVertex)
				hasBeenVisited[visitedVertex] = true
				for _, connection := range rooms.Connection {
					if !hasBeenVisited[connection] {
						toBeVisited = append(toBeVisited, connection)
						toBeVisited = toBeVisited[1:]

					}
				}

			}
		}
	}

	return visited
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
				for _, connection := range rooms.Connection {
					if !hasBeenVisited[connection] {
						visited = append(visited, toBeVisited[len(toBeVisited)-1])
						hasBeenVisited[visitedVertex] = true
						toBeVisited = toBeVisited[:len(toBeVisited)-1]
						toBeVisited = append(toBeVisited, connection)

					}
				}
			}
		}
	}
	return visited
}
