package models

// space for the structs to be written
type Colony struct {
	NoOfAnts int
	Rooms []*Room
	StartRoom string
	EndRoom string
}

type Room struct{
	Name string // A room identifier
	HouseAndCoordinates map[string][]int // A house and the coordinates
	Link [][]string // This stores connections of room to another room
	RoomAndConnectedLinks map[string][]string
}