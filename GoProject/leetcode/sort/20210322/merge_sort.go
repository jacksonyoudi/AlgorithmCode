package main

import "fmt"

func merge_sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2

	left := merge_sort(nums[0:mid])
	right := merge_sort(nums[mid:len(nums)])

	l, r := 0, 0
	ll, lr := len(left), len(right)

	result := make([]int, 0)

	if l < ll && r < lr {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l += 1
		} else {
			result = append(result, right[r])
			r += 1
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	result := merge_sort(arr)
	fmt.Println(result)
}
