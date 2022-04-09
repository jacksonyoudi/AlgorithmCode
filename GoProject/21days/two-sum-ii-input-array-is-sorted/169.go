package two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	left, right := 0, n-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if target == sum {
			return []int{left+1, right+1}
		} else if target < sum {
			right--
			continue
		} else if target > sum {
			left++
			continue
		}
	}

	return []int{-1, -1}

}
