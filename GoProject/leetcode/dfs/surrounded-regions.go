package dfs



// 从边向里面扩张
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	m, n := len(board), len(board[0])

	// 边界上的
	for i := 0; i < m; i++ {
		dfs(board, i, 0)
		dfs(board, i, n-1)
	}

	for i := 0; i < n-1; i++ {
		dfs(board, 0, i)
		dfs(board, m-1, i)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}

// dfs 递归中遍历数据

func dfs(board [][]byte, x, y int) {
	m, n := len(board), len(board[0])


	if x < 0 || y < 0 || x >= m || y >= n || board[x][y] != 'O' {
		return
	}

	board[x][y] = 'A'

	// 方向(1, 0), (-1, 0), (0, 1) (0, -1)
	dfs(board, x+1, y)
	dfs(board, x-1, y)
	dfs(board, x, y+1)
	dfs(board, x, y-1)

}
