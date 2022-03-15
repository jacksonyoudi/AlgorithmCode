package main

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	ans := 0
	for i, ints := range grid {
		for j, value := range ints {
			if value != 0 {
				dfsVaule := dfs(grid, i, j)
				if dfsVaule >= ans {
					ans = dfsVaule
				}

			}
		}
	}
	return ans

}

//  (1,0)  (0,1), (-1, 0) (0, -1)
func dfs(grid [][]int, i int, j int) int {
	l := len(grid)
	r := len(grid[0])

	if i < 0 || j < 0 || i >= l || j >= r || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return 1 + dfs(grid, i+1, j+0) + dfs(grid, i+0, j+1) + dfs(grid, i-1, j-0) + dfs(grid, i-0, j-1)
}
