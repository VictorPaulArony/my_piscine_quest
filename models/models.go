package models

// space for the structs to be written

type Colony struct {
	NoOfAnts int
	Rooms []*Room
	Link [][]string // This stores connections of room to another room
	StartRoom []int
	EndRoom []int
	RoomAndLinks map[string][]string
}

type Room struct{
	Name string // A room identifier
	HouseAndCoordinates map[string][]int // A house and the coordinates
	
}
