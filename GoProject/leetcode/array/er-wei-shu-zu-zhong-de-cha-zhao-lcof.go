package array

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1

	for i < m && j >= 0 {
		tmp := matrix[i][j]
		if tmp == target {
			return true
		} else if tmp > target {
			j--
		} else {
			i++
		}
	}
	return false

}
