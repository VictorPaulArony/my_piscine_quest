package utils

import (
	"fmt"
	"strings"
)

type Ant struct {
	Id       int
	Path     int
	Position int
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
