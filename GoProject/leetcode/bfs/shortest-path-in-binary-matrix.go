package bfs

func shortestPathBinaryMatrix(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	if grid[m-1][n-1] == 1 {
		return -1
	}

	a, b := bfs(grid, 0, 0)
	if b {
		return 1 + a
	} else {
		return -1
	}

}

// 几步到 n-1 n-1
func bfs(grid [][]int, x, y int) (int, bool) {
	// x + 1 , y + 1 x+1 y + 1

	m, n := len(grid), len(grid[0])

	if x >= m || y >= n  || x < 0 || y < 0 || grid[x][y] == 1 {
		return 0, false
	}

	if x == (m-1) && y == (n-1) && grid[x][y] == 0 {
		return 0, true
	}

	grid[x][y] = 1

	v := make([]int, 0)

	// 方向
	directs := [][]int{
		[]int{
			-1, -1,
		},
		{
			0, -1,
		},
		{
			1, 1,
		},
		{
			1, 0,
		},
		{
			1, 1,
		},
		{
			0, 1,
		},
		{
			-1, 1,
		},
		{
			-1, 0,
		},

	}

	for _, direct := range directs {
		p, s := bfs(grid, x+direct[0], y+direct[1])
		if s {
			v = append(v, p)
		}

	}

	if len(v) == 0 {
		return 1, false
	}

	min := v[0]

	for _, i := range v {
		if i < min {
			min = i
		}
	}
	return min + 1, true

}
