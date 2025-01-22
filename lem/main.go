package main

import (
	"bufio"
	"fmt"
	"log"
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
	Paths    []string
	Position int
	Finished bool
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
	currentPath := []string{}              // start from the starting room
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
		}else{
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
func MoveAnts(numAnts int, path []string, g *Graph) {
	moves := []string{}
	ants := make([]*Ant, numAnts)
	for i := range ants {
		// asign each ant the shortest path
		ants[i] = &Ant{
			Id:       i + 1,
			Paths:    path,
			Position: -1, // all ants start in the start room
		}
	}
	currentAnt := 0
	pathId := 0
	for {
		move := []string{}
		allFinished := true

		// Clear room occupancy
		for _, room := range g.Room {
			room.Visited = false
		}

		for _, ant := range ants {
			if ant.Finished {
				continue
			}
			allFinished = false

			if ant.Position == -1 {
				if currentAnt < numAnts {
					ant.Paths = path[1:] // skip first room
					pathId = (pathId + 1) % len(path)
					ant.Position = 0
					currentAnt++
				}
			}

			if ant.Position >= 0 && ant.Position < len(ant.Paths) {
				currentRoom := ant.Paths[ant.Position]
				if !g.Room[currentRoom].Visited {
					g.Room[currentRoom].Visited = true
					move = append(move, fmt.Sprintf("L%d-%s", ant.Id, currentRoom))
					ant.Position++
					if ant.Position >= len(ant.Paths) {
						ant.Finished = true
					}

				}
			}
		}

		if len(move) > 0 {
			moves = append(moves, strings.Join(move, " "))
		}

		if allFinished {
			break
		}

	}
	for _, move := range moves {
		fmt.Println(move)
		continue
	}
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

	MoveAnts(graph.AntCount, path[0], graph)
}
