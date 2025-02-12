package combinations

import (
	"strings"
)

func isInside(pathSearched []string, rooms []string) bool {
	paths := ""
	roomsString := ""
	for _, pathValue := range pathSearched {
		paths += pathValue
	}
	for _, room := range rooms {
		roomsString += room
	}

	return strings.ContainsAny(paths, roomsString)
}

var (
	pathCombinations     [][]string
	possibleCombinations [][][]string
)

func FindCombinations(paths [][]string) [][]string {
	if len(paths) == 1 {
		return paths
	}
	for i := 0; i < len(paths)-1; i++ {
		pathCombinations = nil
		pathCombinations = append(pathCombinations, paths[i])
		uniqueValues := paths[i][1 : len(paths[i])-1]
		for j := i + 1; j < len(paths)-1; j++ {
			pathSearched := paths[j]
			if isInside(pathSearched, uniqueValues) {
				continue
			}
			pathCombinations = append(pathCombinations, pathSearched)
			j++
		}
		possibleCombinations = append(possibleCombinations, pathCombinations)

	}
	return possibleCombinations[0]
}
