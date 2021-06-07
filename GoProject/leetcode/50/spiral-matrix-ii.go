package _0

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	if n == 1 {
		return [][]int{[]int{1}}
	}

	result := make([][]int, 0)
	for i := 0; i < n; i++ {
		tmp := make([]int, 0)

		for j := 0; j < n; j++ {
			tmp = append(tmp, -1)
		}
		result = append(result, tmp)
	}

	directs := [][]int{
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
		[]int{-1, 0},
	}

	i, j := 0, 0

	cnt := 1
	preDirect := 0
	for cnt <= n*n {
		result[i][j] = cnt
		cnt = cnt+ 1

		nextI := i + directs[preDirect][0]
		nextJ := j + directs[preDirect][1]

		if nextI >= 0 && nextJ >= 0 && nextI < n && nextJ < n && result[nextI][nextJ] == -1 {
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
