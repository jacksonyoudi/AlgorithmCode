package main

import "fmt"

func merge_sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	lall := len(nums)
	mid := lall / 2
	left := merge_sort(nums[0:mid])
	right := merge_sort(nums[mid:])

	result := make([]int, 0)

	ll, lr := len(left), len(right)

	i, j := 0, 0

	for i < ll && j < lr {
		if left[i] < right[j] {
			result = append(result, left[i])
			i = i + 1
		} else {
			result = append(result, right[j])
			j = j + 1
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	result := merge_sort(arr)
	fmt.Println(result)
}