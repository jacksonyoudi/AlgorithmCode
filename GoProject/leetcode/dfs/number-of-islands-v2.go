package dfs

// 采用标记的方案， 相同的标记成一个方案
func numIslandsv2(grid [][]byte) int {
	cnt := 0
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])




	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				cnt++
				islandsDfsv2(grid, i, j)
			}
		}
	}
	return cnt

}

func islandsDfsv2(grid [][]byte, x, y int) {
	m, n := len(grid), len(grid[0])

	if x < 0 || y < 0 || x >= m || y >= n || grid[x][y] != '1' {
		return
	}

	grid[x][y] = '0'

	islandsDfsv2(grid, x+1, y)
	islandsDfsv2(grid, x-1, y)
	islandsDfsv2(grid, x, y+1)
	islandsDfsv2(grid, x, y-1)
}
