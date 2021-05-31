package sort

func partition_array(nums [][]int, start int, end int) int {
	pivot := nums[end][0]
	i := start - 1

	for j := start; j < end; j++ {
		if nums[j][0] <= pivot {
			i += 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[end], nums[i+1] = nums[i+1], nums[end]
	return i + 1

}

func quick_sort_array(nums [][]int, start int, end int) {
	if start < end {
		pivot := partition_array(nums, start, end)
		quick_sort_array(nums, start, pivot-1)
		quick_sort_array(nums, pivot+1, end)
	}

}

func merge(intervals [][]int) [][]int {
	// 先排序
	ll := len(intervals) - 1

	quick_sort_array(intervals, 0, len(intervals)-1)

	// 合并
	result := make([][]int, 0)

	currenIndex := 0
	result = append(result, intervals[0])
	currentVal := intervals[currenIndex]

	for i := 1; i <= ll; i++ {
		value := intervals[i]

		// meger
		if value[0] <= currentVal[1] {
			if currentVal[1] < value[1] {
				currentVal[1] = value[1]
			}
			result[currenIndex] = currentVal
		} else {
			currentVal = intervals[i]
			result = append(result, currentVal)
			currenIndex += 1

		}
	}
	return result

}
