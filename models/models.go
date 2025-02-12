package models

// space for the structs to be written
type Colony struct {
	NoOfAnts  int
	Rooms     []*Room
	StartRoom Room
	EndRoom   Room
}

type Room struct {
	Name                  string           // A room identifier
	HouseAndCoordinates   map[string][]int // A house and the coordinates
	Link                  [][]string       // This stores connections of room to another room
	RoomAndConnectedLinks map[string][]string
}

func (c *Colony) Bfs() [][]string {
	toBeVisited := [][]string{}
	validPaths := [][]string{}

	startRoom := []string{c.StartRoom.Name}
	endRoom := c.EndRoom.Name

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

		for _, room := range c.Rooms {
			if room.Name == currentRoom[len(currentRoom)-1] {
				if _, exists := room.RoomAndConnectedLinks[room.Name]; exists {
					for _, connection := range room.RoomAndConnectedLinks[room.Name] {
						matrix := []string{}
						matrix = append(matrix, currentRoom...)
						if !hasBeenVisited[connection] {
							matrix = append(matrix, connection)
							toBeVisited = append(toBeVisited, matrix)
						}
					}
				}
			}
		}
	}
	return validPaths
}
