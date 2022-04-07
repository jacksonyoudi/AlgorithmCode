package range_sum_query_2d_immutable

type NumMatrix struct {
	data   [][]int
	row    int
	col    int
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	row := len(matrix)
	col := len(matrix[0])
	preSum := make([][]int, row+1)
	preSum[0] = make([]int, col+1)

	for i, item := range matrix {
		preSum[i+1] = make([]int, col+1)
		for j, v := range item {
			preSum[i+1][j+1] = preSum[i+1][j] + preSum[i][j+1] - preSum[i][j] + v
		}
	}

	return NumMatrix{
		row:    row,
		col:    col,
		preSum: preSum,
		data:   matrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] + this.preSum[row1][col1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1]
}
