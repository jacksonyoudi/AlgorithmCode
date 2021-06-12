package main

func partition(nums [][]int, start int, end int) int {
	pivot := nums[end][0]
	i := start - 1
	for j := start; j < end; j++ {
		if nums[j][0] < pivot {
			i += 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[end] = nums[end], nums[i+1]
	return i + 1

}

func quick_sort(nums [][]int, start int, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quick_sort(nums, start, pivot-1)
		quick_sort(nums, pivot+1, end)
	}

}

func isCovered(ranges [][]int, left int, right int) bool {
	if len(ranges) == 0 {
		return false
	}

	quick_sort(ranges, 0, len(ranges)-1)

	ll := len(ranges)

	if left > ranges[ll-1][1] || right < ranges[0][0] || right > ranges[ll-1][1] || left < ranges[0][0] {
		return false
	}
	return true

}
