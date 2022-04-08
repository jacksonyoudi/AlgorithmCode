package rotate_image

func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}

	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}
