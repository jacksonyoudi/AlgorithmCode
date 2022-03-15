package main

func findCircleNum(isConnected [][]int) int {
	provinces := len(isConnected)
	visited := make([]bool, provinces)
	cricles := 0

	for i, _ := range isConnected {
		if !visited[i] {
			dfsO(isConnected, visited, provinces, i)
			cricles++
		}
	}

	return cricles
}

func dfsO(isConnected [][]int, visited []bool, provinces int, i int) {
	for j := 0; j < provinces; j++ {
		if isConnected[i][j] == 1 && !visited[j] {
			visited[j] = true
			dfsO(isConnected, visited, provinces, j)
		}
	}
}
