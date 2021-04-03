package main

import "fmt"

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	result := []int{}

	mid := len(nums) / 2

	left := mergeSort(nums[0:mid])
	right := mergeSort(nums[mid:len(nums)])

	ll := len(left)
	lr := len(right)

	l, r := 0, 0
	for l <= ll && r <= lr {
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
	result := mergeSort(arr)
	fmt.Println(result)
}
