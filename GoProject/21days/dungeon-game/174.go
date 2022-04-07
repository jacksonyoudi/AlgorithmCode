package dungeon_game

import "math"

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])

	// 备忘录
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	return dp(dungeon, 0, 0, memo)
}

func dp(grid [][]int, i, j int, memo [][]int) int {
	m := len(grid)
	n := len(grid[0])

	if i == m-1 && j == n-1 {
		if grid[i][j] >= 0 {
			return 1
		} else {
			return -grid[i][j] + 1
		}
	}
	if i == m || j == n {
		return math.MaxInt
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}
	a1 := dp(grid, i, j+1, memo)
	a2 := dp(grid, i+1, j, memo)
	res := a1 - grid[i][j]
	if a1 > a2 {
		res = a2 - grid[i][j]
	}
	if res <= 0 {
		memo[i][j] = 1
	} else {
		memo[i][j] = res
	}

	return memo[i][j]

}
