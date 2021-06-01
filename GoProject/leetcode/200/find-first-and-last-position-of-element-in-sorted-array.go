package _200

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}

	if len(nums) == 0 {
		return []int{-1, -1}
	}


	for i, v := range nums {
		if v == target {
			if result[0] == -1 {
				result[0] = i
			} else {
				result[1] = i
			}
		}
	}

	if result[0] != -1 && result[1] == -1 {
		result[1] = result[0]
	}

	return result
}
