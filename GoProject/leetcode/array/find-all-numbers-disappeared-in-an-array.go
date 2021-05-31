package array

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}

	ans := make([]int, 0)
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return ans
}
