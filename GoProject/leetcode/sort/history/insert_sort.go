package main

func insertSort(nums []int) {
	l := len(nums)
	for i := 1; i < l; i++ {
		current := nums[i]
		j := i - 1
		for j >= 0 && current < nums[j] {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = current
	}
}
