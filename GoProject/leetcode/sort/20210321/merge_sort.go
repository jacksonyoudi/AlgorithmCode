package main

import "fmt"

func merge_sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	l, r := 0, 0

	mid := len(nums) / 2
	left := merge_sort(nums[0:mid])
	right := merge_sort(nums[mid:len(nums)])

	result := make([]int, 0)

	ll, lr := len(left), len(right)

	for l < ll && r < lr {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
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
