package main

func partition(nums []int, start int, end int) int {
	i := start - 1
	pivot := nums[end]

	for j := start; j < end; j++ {
		if nums[j] < pivot {
			i += 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[end] = nums[end], nums[i+1]

	return i + 1
}

func quick_sort(nums []int, start int, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quick_sort(nums, start, pivot-1)
		quick_sort(nums, pivot+1, end)

	}

}

func getLeastNumbers(arr []int, k int) []int {
	quick_sort(arr, 0, len(arr)-1)
	return arr[:k]

}
