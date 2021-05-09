package dfs

// 采用标记的方案， 相同的标记成一个方案
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	gridmap := [][]int{}
	for i := 0; i < m; i++ {
		a := []int{}
		for j := 0; j < n; j++ {
			a = append(a, 0)
		}
		gridmap = append(gridmap, a)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			i2 := i*300 + j + 311

			islandsDfs(grid, i, j, i2, gridmap)
		}
	}

	m2 := make(map[int]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := gridmap[i][j]
			if k == 0 {
				continue
			}
			if v, ok := m2[k]; !ok {
				m2[k] = v
			}
		}
	}
	return len(m2)

}

func islandsDfs(grid [][]byte, x, y int, su int, gridmap [][]int) {
	m, n := len(grid), len(grid[0])

	if x < 0 || y < 0 || x >= m || y >= n || grid[x][y] != '1' {
		return
	}

	gridmap[x][y] = su
	grid[x][y] = '2'

	islandsDfs(grid, x+1, y, su, gridmap)
	islandsDfs(grid, x-1, y, su, gridmap)
	islandsDfs(grid, x, y+1, su, gridmap)
	islandsDfs(grid, x, y-1, su, gridmap)
}
