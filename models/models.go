package models

// space for the structs to be written

type Colony struct {
	NoOfAnts int
	Rooms *[]Room
}

type Room struct{
	Name string
	Coordinates map[string][]int
	Link map[string]int
}
