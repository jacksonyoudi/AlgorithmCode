package _200

func rotate(matrix [][]int) {
	n := len(matrix)
	if n <= 1 {
		return
	}

	x := n / 2
	y := (n + 1) / 2

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}

}
