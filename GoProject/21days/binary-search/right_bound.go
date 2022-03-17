package binary_search

// 查找最右边的那个
func right_bound(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left != 0 && nums[left-1] == target {
		return left - 1
	} else {
		return -1
	}

}
