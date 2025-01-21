package main

import "fmt"

// Graph structure
type Graph struct {
	Node map[int][]int
}

func main() {
	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(4, 6)
	g.AddEdge(5, 6)

	start := 0
	target := 6
	if path, found := g.BFS(start, target); found {
		fmt.Printf("Shortest path from %d to %d: %v\n", start, target, path)
	} else {
		fmt.Printf("No path found from %d to %d\n", start, target)
	}
}

// NewGraph creates a new graph
func NewGraph() *Graph {
	return &Graph{Node: make(map[int][]int)}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(v, w int) {
	g.Node[v] = append(g.Node[v], w)
	g.Node[w] = append(g.Node[w], v) // For undirected graph
}


// function to show implementation of bfs algorithm in go
func (g *Graph)BFS(start int, end int) ([]int, bool) {
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
			return Oder(paths, start, end), true
		}

		for _, neighbor := range g.Node[currentNode] {
			if !visited[neighbor] {
				visited[neighbor] = true
				paths[neighbor] = currentNode // track the
				queue = append(queue, neighbor)
			}
		}
	}

	return nil, false //no path found
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
