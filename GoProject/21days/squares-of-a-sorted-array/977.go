package squares_of_a_sorted_array

func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	newNums := make([]int, len(nums))

	pre := nums[0] * nums[0]

	idx := 0
	for i, v := range nums {
		newNums[i] = v * v

		if pre > newNums[i] {
			idx = i
		}
		pre = newNums[i]
	}
	if idx == 0 {
		return newNums
	} else if idx == len(newNums)-1 {
		result := make([]int, len(newNums))
		for i, v := range newNums {
			result[len(newNums)-i-1] = v
		}
		return result

	} else {
		n := len(newNums)
		result := make([]int, 0)

		i, j := idx-1, idx
		cnt := 0

		// 归并排序
		for i >= 0 && j < n && cnt <= n {
			if newNums[i] <= newNums[j] {
				result = append(result, newNums[i])
				i--
			} else {
				result = append(result, newNums[j])
				j++
			}
			cnt++
		}
		if i > 0 {
			result = append(result, newNums[0:i]...)
		}

		if j < n {
			result = append(result, newNums[j:]...)
		}

		return result
	}

}
