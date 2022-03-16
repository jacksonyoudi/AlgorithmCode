package binary_search

func searchHelper(nums []int, start, end int, target int) int {
	if start > end {
		return -1
	}
	mid := start + (end-start)/2
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return searchHelper(nums, mid+1, end, target)
	} else if nums[mid] > target {
		return searchHelper(nums, start, mid-1, target)
	}
	return -1
}

func search2(nums []int, target int) int {
	return searchHelper(nums, 0, len(nums)-1, target)
}
