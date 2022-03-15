package main

import "fmt"

func partition_2(nums []int, start int, end int) int {
	pivot := nums[end]
	i := start - 1

	for j := start; j < end; j++ {
		if nums[j] < pivot {
			i += 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[end] = nums[end], nums[i+1]
	return i + 1

}

func quickSort(nums []int, start int, end int) {
	if start < end {
		pivot := partition_2(nums, start, end)
		quickSort(nums, start, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
