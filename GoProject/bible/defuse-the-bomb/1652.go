package defuse_the_bomb

func helper_a(nums []int, index int, k int) int {
	l := len(nums)
	r := 0
	for i := 0; i < k; i++ {
		if index == l {
			index = 0
		}
		r += nums[index]
		index++
	}
	return r
}

func helper_b(nums []int, index int, k int) int {
	l := len(nums)
	r := 0
	for i := 0; i < k; i++ {
		if index == -1 {
			index = l - 1
		}
		r += nums[index]
		index--
	}
	return r
}

func decrypt(code []int, k int) []int {
	l := len(code)
	result := make([]int, l)
	if k == 0 {
		for i := 0; i < l; i++ {
			result[i] = 0
		}
	} else if k > 0 {
		for i := 0; i < l; i++ {
			result[i] = helper_a(code, i+1, k)
		}
	} else {
		for i := 0; i < l; i++ {
			result[i] = helper_b(code, i-1, k)
		}

	}
	return result

}
