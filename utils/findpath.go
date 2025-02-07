package utils

// function to find the paths from start to end of the graph
func (g *Graph) FindPaths(start, end string) [][]string {
	allPaths := [][]string{}
	path := []string{}
	visited := make(map[string]bool)

	// DFS function to find the all the paths
	var dfs func(node string)
	dfs = func(node string) {
		visited[node] = true
		path = append(path, node)
		if node == end {
			formedPath := make([]string, len(path))
			copy(formedPath, path)
			allPaths = append(allPaths, formedPath)
		} else {
			for _, nextNode := range g.Rooms[node].Links {
				if !visited[nextNode] {
					dfs(nextNode)

					// all for the backdracking of to go to other rooms
					visited[nextNode] = false

					// Backtrack: remove node from path
					path = path[:len(path)-1]
				}
			}
		}
	}
	// the function is called recusivelly(if it reaches at the end the function is called again)
	dfs(start)

	// Sort the paths based on length
	// sort.Slice(allPaths, func(i, j int) bool {
	// 	return len(allPaths[i]) < len(allPaths[j])
	// })
	allPaths = SortPath(allPaths)

	return allPaths
}

// function to sort the paths from the shortest to the longest
func SortPath(arr [][]string) [][]string {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if len(arr[j]) > len(arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
