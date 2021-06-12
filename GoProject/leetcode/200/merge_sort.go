package _200

func merge_sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	l := merge_sort(nums[0:mid])
	r := merge_sort(nums[mid:])

	ll, lr := len(l), len(r)

	i, j := 0, 0

	result := make([]int, 0)

	for i < ll && j < lr {
		if l[i] < r[j] {
			result = append(result, r[j])
			j++
		} else {
			result = append(result, l[i])
			i++
		}
	}

	result = append(result, l[i:]...)
	result = append(result, r[j:]...)

	return result

}
