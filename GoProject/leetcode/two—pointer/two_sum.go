package main

import "fmt"

func twoSum(nums []int, target int) []int {
	start := 0
	end := len(nums) - 1

	result := []int{}
	for start < end {
		if nums[start]+nums[end] == target {
			result = append(result, start)
			result = append(result, end)
			start++
			end--
		} else if nums[start]+nums[end] < target {
			start++
		} else {
			end--
		}
	}
	return result

}

func main() {
	fmt.Println("hello")
}
