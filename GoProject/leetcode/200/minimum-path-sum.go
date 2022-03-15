package _200

func minPathSumDfs(grid [][]int, n, m, x, y int, path int) int {
	if x == n-1 || y == m-1 {
		return path + grid[x][y]
	}

	if x >= n || y >= m {
		return -1
	}

	a := minPathSumDfs(grid, n, m, x+1, y, path)
	b := minPathSumDfs(grid, n, m, x, y+1, path)
	if a == -1 && b == -1 {
		return -1
	}

	if a == -1 && b != -1 {
		return b + grid[x][y]
	}

	if a != -1 && b == -1 {
		return a + grid[x][y]
	}

	if a > b {
		return b + grid[x][y]
	} else {
		return a + grid[x][y]
	}
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	return minPathSumDfs(grid, len(grid), len(grid[0]), 0, 0, 0)
}
