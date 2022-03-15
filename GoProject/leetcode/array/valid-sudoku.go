package array

func isOK(digits []byte) bool {
	m := make(map[byte]int)

	for _, digit := range digits {
		if digit == '.' {
			continue
		}
		_, ok := m[digit]
		if ok {
			return false
		} else {
			m[digit] = 1
		}

	}
	return true

}

func isValidSudoku(board [][]byte) bool {
	// row
	for _, row := range board {
		if !isOK(row) {
			return false
		}
	}

	// lie
	for i := 0; i < 9; i++ {
		a := []byte{}

		for j := 0; j < 9; j++ {
			a = append(a, board[j][i])
		}
		if !isOK(a) {
			return false
		}
	}

	// 3 * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			a := []byte{}
			for ii := i * 3; ii < i*3+3; ii++ {
				for jj := j * 3; jj < j*3+3; jj++ {
					a = append(a, board[ii][jj])
				}
			}

			if !isOK(a) {
				return false
			}
		}

	}

	return true

}
