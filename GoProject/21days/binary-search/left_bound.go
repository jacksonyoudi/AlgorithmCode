package binary_search

func left_bound(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (left-right)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}
