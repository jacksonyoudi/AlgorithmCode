package sort

func merge_sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2

	l, r := 0, 0

	left := merge_sort(nums[0:mid])
	right := merge_sort(nums[mid:])

	result := make([]int, 0)
	ll := len(left)
	lr := len(right)

	for l < ll && r < lr {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	if l < ll {
		result = append(result, left[l:]...)
	}
	if r < lr {
		result = append(result, right[r:]...)
	}

	return result
}



