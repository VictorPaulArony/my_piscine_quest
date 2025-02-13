package utils

import "fmt"

// function to print the output of the program
func Printout(fileName string) {
	colony := FileReader(fileName)
	if colony == nil {
		return
	}

	// Print number of ants
	fmt.Println(colony.NoOfAnts)

	// Print rooms with start and end markers in correct positions
	var startPrinted bool
	var endPrinted bool

	// Print start room first
	fmt.Println("##start")
	for name, coord := range colony.StartRoom.HouseAndCoordinates {
		fmt.Printf("%s %d %d\n", name, coord[0], coord[1])
		if startPrinted{ 
		break
		}
	}

	// Print middle rooms
	for _, room := range colony.Rooms {
		// Skip start and end rooms
		if room.Name == colony.StartRoom.Name || room.Name == colony.EndRoom.Name {
			continue
		}
		for name, coord := range room.HouseAndCoordinates {
			fmt.Printf("%s %d %d\n", name, coord[0], coord[1])
		}
	}

	// Print end room last
	fmt.Println("##end")
	for name, coord := range colony.EndRoom.HouseAndCoordinates {
		fmt.Printf("%s %d %d\n", name, coord[0], coord[1])
		if endPrinted {
		break
		}
	}

	// Print links
	for _, room := range colony.Rooms {
		for _, v := range room.Link {
			fmt.Printf("%s-%s\n", v[0], v[1])
		}
		if !endPrinted {
			break
			}
	}
	fmt.Println()

	// Process paths and ant movements
	startRoom := colony.StartRoom
	endRoom := colony.EndRoom
	totalNumber := colony.NoOfAnts
	allPath := colony.Bfs()
	bestPaths := FilterPath(allPath, startRoom.Name, endRoom.Name)
	antDistribution := AntDistribution(bestPaths, totalNumber)
	MoveAnts(antDistribution, bestPaths)
}
