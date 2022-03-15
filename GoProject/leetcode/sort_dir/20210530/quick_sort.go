package main

import "fmt"

func partition(nums []int, start int, end int) int {
	i := start - 1
	pivot := nums[end]

	for j := start; j < end; j++ {
		if nums[j] < pivot {
			i++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}

	nums[end], nums[i+1] = nums[i+1], nums[end]

	return i + 1
}

func quick_Sort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quick_Sort(nums, start, pivot-1)
		quick_Sort(nums, pivot+1, end)
	}

}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	quick_Sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
