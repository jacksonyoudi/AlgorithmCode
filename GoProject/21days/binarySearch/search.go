package binarySearch

func binary_search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			return mid
		}
	}

	return -1
}

func left_bound(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 别返回，锁定左侧边界
			right = mid - 1
		}

	}

	// 最后要检查 left 越界的情况
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left

}

func right_bound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 别返回，锁定左侧边界
			left = mid + 1
		}

	}

	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
