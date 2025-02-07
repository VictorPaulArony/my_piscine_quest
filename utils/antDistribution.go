package utils

import "math"

// function to distribute the ants to the best paths
func AntDistribution(paths [][]string, numAnts int) [][]int {
	// initialize the array to hold the distibution
	distribution := make([][]int, len(paths))

	// calculate the length of each path excluding the start node
	pathlength := make([]int, len(paths))
	for i, path := range paths {
		pathlength[i] = len(path) - 1
		// the length of the path is number of m=nodes minus one
	}

	// iterate through all the ants to assing them path
	for i := 1; i <= numAnts; i++ {
		bestPath := 0 // the path with the shortest arival time
		bestTime := math.MaxInt64

		// now iterate through all the paths to determine the shortest
		for i := range paths {
			// calc the arrival time for this path
			arrivalTime := len(distribution[i]) + pathlength[i]
			if arrivalTime < bestTime {
				bestPath = i
				bestTime = arrivalTime
			}
		}
		distribution[bestPath] = append(distribution[bestPath], i)
	}
	return distribution
}
