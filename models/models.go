package models

type Room struct {
	RoomNumber int
	Coordinate []int
	Connection []int
}
type Colony struct {
	NumAnts int
	Rooms   map[string]*Room
	Start   Room
	End     Room
}

