package main

import "fmt"

// type graph to store the nodes
type Graph struct {
	Nodes map[int][]int
}

// function to initialaize the graph
func NewGraph() *Graph {
	return &Graph{Nodes: make(map[int][]int)}
}

// function to populate the graph values for the node
func (g *Graph) AddNodes(x, y int) {
	g.Nodes[x] = append(g.Nodes[x], y)
	g.Nodes[y] = append(g.Nodes[y], x)
}

// function to perform the recusion to find all the posible paths
func (g *Graph) DFS(start, target int) [][]int {
	var path []int
	visited := make(map[int]bool)
	var allPaths [][]int

	// helper function to perform DFS recursively
	var dfsHelper func(node int)
	dfsHelper = func(node int) {
		// mark the current node ass visted and add to path
		visited[node] = true
		path = append(path, node)

		// if the cuurent node is the target save all the [paths]
		if node == target {
			// append a copy of the path to all the paths
			newPath := make([]int, len(path))
			copy(newPath, path)
			allPaths = append(allPaths, newPath)
		} else {
			// loop recusively through all nodes
			for _, neighbor := range g.Nodes[node] {
				if !visited[neighbor] {
					dfsHelper(neighbor)
				}
			}
		}
		// Backtrack: unmark the current node and remove it from the path
		visited[node] = false
		path = path[:len(path)-1]
	}
	dfsHelper(start)

	return allPaths
}

func main() {
	g := NewGraph()
	g.AddNodes(0, 1)
	g.AddNodes(0, 2)
	g.AddNodes(1, 3)
	g.AddNodes(2, 3)
	g.AddNodes(3, 4)

	start := 0
	target := 4
	paths := g.DFS(start, target)

	fmt.Printf("All paths from %d to %d:\n", start, target)
	for _, path := range paths {
		fmt.Println(path)
	}
}
