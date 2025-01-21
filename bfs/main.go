package main

import "fmt"

func main() {
	graph := map[int][]int{
		5: {3, 7},
		3: {2, 4},
		7: {8},
		2: {},
		4: {8},
		8: {},
	}
	found := BFS(graph, 5)
	fmt.Println(found)
}

// function to show implementation of bfs algorithm in go
func BFS(graph map[int][]int, start int) []int {
	// store all the paths visited
	paths := []int{}
	visited := make(map[int]bool) // to mark the visited node

	// queue to track the nodes
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		// remove the first node to add the next
		currentNode := queue[0]
		queue = queue[1:]

		// store all the paths visted
		paths = append(paths, currentNode)
		for _, neighbor := range graph[currentNode] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return paths
}
