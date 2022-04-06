package car_pooling

func carPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1001)
	result := make([]int, 1001)

	for _, trip := range trips {
		val, start, end := trip[0], trip[1], trip[2]
		nums[start] += val
		nums[end] -= val
	}

	for i, v := range nums {
		if i == 0 {
			result[i] = nums[i]
		} else {
			result[i] = v + result[i-1]
		}

		if result[i] > capacity {
			return false
		}
	}

	return true
}
