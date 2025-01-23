package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Room represents a room in the graph with its name, coordinates, visited status, and links to other rooms.
type Room struct {
	Name    string   // Name of the room
	x, y    int      // Coordinates of the room
	Visited bool     // Indicates if the room has been visited
	Link    []string // Links to other rooms
}

// Graph represents the entire structure of the ant colony, including rooms and the number of ants.
type Graph struct {
	AntCount int              // Number of ants
	Room     map[string]*Room // Map of room names to Room objects
	Start    Room             // Starting room
	End      Room             // Ending room
}

type Ant struct {
	Id       int
	Path     int
	Position int
}

// ReadFile reads the input file and constructs the graph structure.
func ReadFile(fileName string) *Graph {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error while reading the file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Initialize the graph
	g := &Graph{Room: make(map[string]*Room)}

	// Check if the file is empty
	if !scanner.Scan() {
		log.Println("Empty file")
		return nil
	}

	// Read the number of ants
	totalAnts, err := strconv.Atoi(scanner.Text())
	if err != nil || totalAnts <= 0 {
		log.Println("Invalid number of ants")
		return nil
	}
	g.AntCount = totalAnts

	// Flags for start and end
	start := false
	target := false

	// Process each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue // Skip empty lines
		}

		// Check for special commands
		if line == "##start" {
			start = true
			continue
		} else if line == "##end" {
			target = true
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue // Skip comments
		}

		// Handle room definitions
		if strings.Contains(line, " ") {
			parts := strings.Fields(line)
			name := parts[0]
			x, err1 := strconv.Atoi(parts[1])
			y, err2 := strconv.Atoi(parts[2])
			if err1 != nil || err2 != nil {
				fmt.Printf("Invalid coordinates for room: %s\n", name)
				return nil
			}

			room := &Room{
				Name: name,
				x:    x,
				y:    y,
			}
			if start {
				g.Start = *room
				start = false
			} else if target {
				g.End = *room
				target = false
			}
			g.Room[name] = room

			// Handle tunnel connections
		} else if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				fmt.Println("ERROR: Invalid data format, cannot link a room to itself")
				return nil
			}

			room1, exist1 := g.Room[parts[0]]
			room2, exist2 := g.Room[parts[1]]

			if !exist1 || !exist2 {
				fmt.Println("Link references non-existent room")
				return nil
			}
			room1.Link = append(room1.Link, parts[1])
			room2.Link = append(room2.Link, parts[0]) // Create bidirectional link
		}
	}

	return g
}

// function to fined the shortest path
func FindPath(g *Graph, start, target string) [][]string {
	paths := [][]string{}
	currentPath := []string{}        // start from the starting room
	visited := make(map[string]bool) // to avoid revisting the rooms

	// define a helper function for the DFS
	var dfs func(node string)
	dfs = func(node string) {
		visited[node] = true
		currentPath = append(currentPath, node)

		if node == target {
			newPath := make([]string, len(currentPath))
			copy(newPath, currentPath)
			paths = append(paths, newPath)
		} else {
			// Explore all linked rooms
			for _, neighbor := range g.Room[node].Link {
				if !visited[neighbor] {
					dfs(neighbor)
				}
			}
		}

		// Backtrack: unmark the current node and remove it from the path
		visited[node] = false
		currentPath = currentPath[:len(currentPath)-1]
	}
	dfs(start)
	return paths
}

// function to move the ants in the rooms
func MoveAnts(antDistribution [][]int, paths [][]string) {
	// initialize ant position based on the distribution
	var ants []Ant
	for pathId, allAnts := range antDistribution {
		for _, ant := range allAnts {
			ants = append(ants, Ant{ant, pathId, 0})
		}
	}


	// simulate the movements until all ants reach end point
	for len(ants) > 0 {
		moves := []string{}
		var allAnts []Ant
		visited := make(map[string]bool)
		for _, ant := range ants {
			if ant.Position < len(paths[ant.Path])-1 {
				currentRoom := paths[ant.Path][ant.Position]
				nextRoom := paths[ant.Path][ant.Position+1]
				link := currentRoom + "-" + nextRoom

				if !visited[link] {
					moves = append(moves, fmt.Sprintf("L%d-%s", ant.Id, nextRoom))
					allAnts = append(allAnts, Ant{ant.Id, ant.Path, ant.Position + 1})
					visited[link] = true
				} else {
					// othewise keep the ant in its currect position
					allAnts = append(allAnts, ant)
				}
			}
		}
		// print moves for the current turn
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
		// update position for the next turn
		ants = allAnts
		
	}
}

// function that place the ants in the paths optimamlly
func AntPlacer(paths [][]string, numAnts int) [][]int {
	// initialize an array to hold the distribution of ants across the paths
	distribution := make([][]int, len(paths))

	// calculate the length of each path excluding the start node
	pathLengths := make([]int, len(paths))
	for i, path := range paths {
		pathLengths[i] = len(path) - 1
		// the length of the path is number of m=nodes minus one
	}

	// iterate through each ant to assign them to the paths
	for ant := 1; ant <= numAnts; ant++ {
		bestPath := 0 // path with the shortest arival time
		bestTime := math.MaxInt64//

		// iterate through all the paths to determine the shortest
		for i:= range paths {
			// calc the arival time for this path
			// current ants on the path = len(distrbution[i])
			arivalTime := len(distribution[i]) + pathLengths[i]
			// updating the best path if the path arrival time is shorter
			if arivalTime < bestTime {
				bestPath = i
				bestTime = arivalTime
			}
		}
		// assign the ants to the best path by appending it to thr corresponding distribution slice
		distribution[bestPath] = append(distribution[bestPath], ant)
	}
	fmt.Println(distribution)
	return distribution
}

// main function to execute the program
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}

	// Read the graph from the input file
	graph := ReadFile(os.Args[1])

	path := FindPath(graph, graph.Start.Name, graph.End.Name)
	fmt.Println("Shortest path from start to end:")
	fmt.Println(path)

	antCount := AntPlacer(path, graph.AntCount)
	MoveAnts(antCount, path)
}
