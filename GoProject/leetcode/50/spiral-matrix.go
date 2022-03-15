package _0

import "math"

func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0)

	m := len(matrix)
	if m == 0 {
		return result
	}

	n := len(matrix[0])
	if n == 0 {
		return result
	}

	i, j := 0, 0

	directs := [][]int{
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
		[]int{-1, 0},
	}

	cnt := 0
	preDirect := 0
	for cnt < m*n {
		result = append(result, matrix[i][j])
		matrix[i][j] = math.MinInt32
		cnt++

		nextI := i + directs[preDirect][0]
		nextJ := j + directs[preDirect][1]

		if nextI >= 0 && nextJ >= 0 && nextI < m && nextJ < n && matrix[nextI][nextJ] != math.MinInt32 {
			i = nextI
			j = nextJ
		} else {
			preDirect = preDirect + 1
			if preDirect == 4 {
				preDirect = 0
			}
			i = i + directs[preDirect][0]
			j = j + directs[preDirect][1]
		}
	}

	return result

}
