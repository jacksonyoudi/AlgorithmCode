package zero_matrix_lcci

func helper(matrix [][]int, x, y int) {
	for i, nums := range matrix {
		if i == x {
			for j, _ := range nums {
				matrix[i][j] = 0
			}
		} else {
			matrix[i][y] = 0
		}
	}
}

func setZeroes(matrix [][]int) {
	tmp := make([][]int, 0)

	for i, nums := range matrix {
		for j, v := range nums {
			if v == 0 {
				tmp = append(tmp, []int{i, j})
			}
		}
	}

	for _, item := range tmp {
		helper(matrix, item[0], item[1])
	}

}
