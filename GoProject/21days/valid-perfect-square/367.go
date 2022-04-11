package valid_perfect_square

func isPerfectSquare(num int) bool {
	left, right := 1, num

	for left <= right {
		mid := left + (right-right)>>1
		if mid*mid < num {
			left = mid + 1
		} else if mid*mid > num {
			right = mid - 1
		} else if mid*mid == num {
			return true
		}
	}

	return false
}
