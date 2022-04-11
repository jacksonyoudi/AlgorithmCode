package peak_index_in_a_mountain_array

func peakIndexInMountainArray(arr []int) int {
	// 二分法
	left, right := 1, len(arr)-2
	for left <= right {
		mid := left + (right-left)>>1
		// case-1, 中值大于左右两边值
		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return mid
			// case-2，中值小于左边，大值出现在左边
		} else if arr[mid] < arr[mid-1] {
			right = mid - 1
			// case-3，中值小于右边，大值在右边
		} else {
			left = mid + 1
		}
	}
	return left
}

func peakIndexInMountainArray2(nums []int) int {
	// 双指针，左边严格单增，峰顶右边单减，峰顶不能是两边的边值
	// 头指针右移，尾指针左移，看是否同个元素
	head, tail := 0, len(nums)-1
	// 头指针右移
	for head+1 < len(nums) && nums[head] < nums[head+1] {
		head++
	}
	// 尾指针左移
	for tail > 0 && nums[tail] < nums[tail-1] {
		tail--
	}
	// head == tail 且去头尾边值
	if head > 0 && tail < len(nums)-1 && head == tail {
		return head
	}
	return -1
}
