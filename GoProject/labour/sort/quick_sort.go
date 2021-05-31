package sort

import "fmt"

func partition(nums []int, start int, end int) int {
	pivot := nums[end]
	i := start - 1

	for j := start; j < end; j++ {
		if nums[j] <= pivot {
			i += 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[end] = nums[end], nums[i+1]
	return i + 1

}

func quick_sort(nums []int, start int, end int) {
	if start < end {
		par := partition(nums, start, end)
		quick_sort(nums, start, par-1)
		quick_sort(nums, par+1, par)
	}

	fmt.Println(nums)
}
