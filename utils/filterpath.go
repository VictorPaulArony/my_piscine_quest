package utils

// function that filters paths and returns the best solution without any conflicts
func FilterPath(AllPaths [][]string, start string, end string) [][]string {
	BestSolution := [][]string{}

	// Iterate through all paths to find the best solution
	for i := 0; i < len(AllPaths); i++ {
		CurrentSolution := [][]string{AllPaths[i]}
		for j := 0; j < len(AllPaths); j++ {
			// Ensure paths are not compared with themselves
			if i != j && CheckPath(CurrentSolution, AllPaths[j], start, end) {
				CurrentSolution = append(CurrentSolution, AllPaths[j])
			}
		}

		// Update the best solution if the current one is found to be longer 
		if len(CurrentSolution) > len(BestSolution) {
			BestSolution = CurrentSolution
		}
	}
	
	// Return the best solution found
	return BestSolution
}

// CheckPath checks if the current path can be added to the existing path without repeating the rooms to avoid resourse queue(waiting)
func CheckPath(path [][]string, current []string, start string, end string) bool {
	// Iterate through the current path to check for conflicts
	for i := 0; i < len(path); i++ {
		for _, room := range path[i] {
			// Skip start and end rooms
			if room == start || room == end {
				continue
			}

			// Check if the room already exists in the current path (conflict)
			for _, curRoom := range current {
				if curRoom == room && curRoom != start && curRoom != end {
					return false 
				}
			}
		}
	}
	return true
}
