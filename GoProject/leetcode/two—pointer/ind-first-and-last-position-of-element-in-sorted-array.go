package main

import "fmt"

func searchRange(nums []int, target int) []int {
	start, end := -1, -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			start = i
			break
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == target {
			fmt.Println(i)
			end = i
			break
		}
	}

	return []int{start, end}

}

func main() {
	fmt.Println(searchRange([]int{1}, 1))

}
