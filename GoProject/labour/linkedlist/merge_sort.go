package linkedlist

func merge_sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2

	left := merge_sort(nums[0:mid])
	right := merge_sort(nums[mid:])

	ll := len(left)
	lr := len(right)

	l, r := 0, 0

	result := make([]int, 0)

	for l < ll && r < lr {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l += 1
		}

		if left[l] > right[r] {
			result = append(result, right[r])
			r += 1
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
