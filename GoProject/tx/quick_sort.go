package tx

func partition(nums []int, start, end int) int {
	pivot := nums[end]
	i := start - 1

	for j := start; j < end; j++ {
		if nums[j] < pivot {
			i += 1
			nums[j], nums[i] = nums[i], nums[j]
		}

	}

	nums[i+1], nums[end] = nums[end], nums[i+1]
	return i + 1
}

func quick_sort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)

		quick_sort(nums, start, pivot-1)
		quick_sort(nums, pivot+1, end)

	}

}
