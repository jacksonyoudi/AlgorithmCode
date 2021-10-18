package tx

func merge_sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2

	left := nums[0:mid]
	right := nums[mid:]
	ll := len(left)
	lr := len(right)

	result := make([]int, 0)
	i, j := 0, 0

	for i < ll && j < lr {
		if left[i] < right[j] {

		} else {

		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result

}
