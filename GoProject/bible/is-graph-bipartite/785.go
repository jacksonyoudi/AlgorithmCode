package is_graph_bipartite

func bfs(graph [][]int, node int, color int, visited map[int]bool, colors map[int]int) bool {
	visited[node] = true

	if b, ok := colors[node]; !ok {
		colors[node] = color
	} else {
		color = b
	}

	for _, children := range graph[node] {
		if _, ok := visited[children]; ok {
			if c, ok := colors[children]; ok {
				if c != -color {
					return false
				}
			} else {
				colors[children] = -color
			}
		} else {
			return bfs(graph, children, -color, visited, colors)
		}
	}

	return true

}

func isBipartite(graph [][]int) bool {
	visited := make(map[int]bool)
	colors := make(map[int]int)

	for i, _ := range graph {
		ok := bfs(graph, i, 1, visited, colors)
		if !ok {
			return false
		}
	}

	return true
}
