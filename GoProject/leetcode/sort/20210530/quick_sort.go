package main

import "fmt"

func partition(nums []int, start, end int) int {
	i := start - 1
	pivot := nums[end]

	for j := start; j < end; j++ {
		if nums[j] < pivot {
			i++
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


func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	quick_sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}