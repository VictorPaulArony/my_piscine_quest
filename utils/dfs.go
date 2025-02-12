package utils

// Takes the graph, the start and the end and returns the list of all valid paths.
func Dfs(start, end string, graph map[string][]string, visited map[string]bool, path []string) [][]string {
	var possiblePaths [][]string
	// var path []string

	path = append(path, start) // add the current node in the path
	if start == end {
		tmp := make([]string, len(path))
		copy(tmp, path)
		return [][]string{tmp}
	}

	visited[start] = true // mark the node as visited

	for _, neighbor := range graph[start] {
		if !visited[neighbor] {
			paths := Dfs(neighbor, end, graph, visited, path)
			possiblePaths = append(possiblePaths, paths...)
		}
	}
	visited[start] = false
	
	return possiblePaths
}