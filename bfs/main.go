package main

import "fmt"

func main() {
	graph := map[int][]int{
		0: {0, 1},
		1: {0, 2},
		2: {1, 3},
		3: {1, 4},
		4: {2, 5},
		5: {3, 6},
		6: {4, 6},
		7: {5, 6},
	}
	start := 0
	target := 6
	path := BFS(graph, start, target)
	fmt.Printf("Shortest path from %d to %d: %v\n", start, target, path)
}

// function to show implementation of bfs algorithm in go
func BFS(graph map[int][]int, start int, end int) []int {
	visited := make(map[int]bool) // to mark the visited node
	// queue to track the nodes
	queue := []int{start}
	visited[start] = true
	// store all the paths visited\
	paths := make(map[int]int)

	// BFS loop
	for len(queue) > 0 {
		// remove the first node to add the next
		currentNode := queue[0]
		queue = queue[1:]

		if currentNode == end {
			return Oder(paths, start, end)
		}

		for _, neighbor := range graph[currentNode] {
			if !visited[neighbor] {
				visited[neighbor] = true
				paths[neighbor] = currentNode // track the
				queue = append(queue, neighbor)
			}
		}
	}

	return nil
}

// function to orderthe path from start to end
func Oder(paths map[int]int, start, end int) []int {
	path := []int{}
	for node := end; node != start; node = paths[node] {
		path = append([]int{node}, path...)
	}
	path = append([]int{start}, path...)
	return path
}
