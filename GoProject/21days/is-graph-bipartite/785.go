package is_graph_bipartite

func isBipartite(graph [][]int) bool {
	n := len(graph)
	data := &Data{
		Colors:  make([]int, n),
		Visited: make([]int, n),
	}

	return bfs(graph, data, 0, -1, n)

}

func bfs(graph [][]int, data *Data, index int, color, n int) bool {
	if index >= n {
		return true
	}
	data.Visited[index] = 1
	nodes := graph[index]
	for _, node := range nodes {
		if data.Visited[node] == 1 {
			if data.Colors[node] == color {
				return false
			} else {
				continue
			}
		}
	}
	return bfs(graph, data, index+1, -color, n)

}

type Data struct {
	Colors  []int
	Visited []int
}
