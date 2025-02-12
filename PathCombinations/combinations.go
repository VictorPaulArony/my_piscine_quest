package combinations

import (
	"strconv"
	"strings"
)

func isInside(pathSearched []int, rooms []int) bool {
	paths := ""
	roomsString := ""
	for _, pathValue := range pathSearched {
		paths += strconv.Itoa(pathValue)
	}
	for _, room := range rooms {
		roomsString += strconv.Itoa(room)
	}

	return strings.ContainsAny(paths, roomsString)
}

var (
	pathCombinations     [][]int
	possibleCombinations [][][]int
)

func FindCombinations(paths [][]int) [][]int {
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
