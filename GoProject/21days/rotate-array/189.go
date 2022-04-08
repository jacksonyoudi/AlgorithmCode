package rotate_array

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	result := make([]int, n)

	for i, v := range nums {
		index := (i + k) % n
		result[index] = v
	}
	for i, v := range result {
		nums[i] = v
	}
}
