package utils

import (
	"fmt"
)

func BfsShortestPath(start, end string, graph map[string][]string) []string {
	fmt.Println("This is Breadth First Search")
	queue := [][]string{{start}}      // This is the queue to store the unexplored nodes
	explored := make(map[string]bool) // A set to store the explored vertices
	for len(queue) > 0 {
		// Dequeue the first path from the queue
		path := queue[0]
		queue = queue[1:]

		// Get the last node in the path
		node := path[len(path)-1]
		if node == end {
			return path
		}

		// Mark as visited
		if !explored[node] {
			explored[node] = true

			// Add all unvisited neighbors in the queue
			for _, neighbor := range graph[node] {
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return nil
}
