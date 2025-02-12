package models

// space for the structs to be written
type Colony struct {
	NoOfAnts int
	Rooms []*Room
	StartRoom Room
	EndRoom Room
}

type Room struct{
	Name string // A room identifier
	HouseAndCoordinates map[string][]int // A house and the coordinates
	Link [][]string // This stores connections of room to another room
	RoomAndLinks map[string][]string
}