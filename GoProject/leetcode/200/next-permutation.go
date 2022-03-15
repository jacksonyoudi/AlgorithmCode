package _200


func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2

	// 找出第一个 降序的 位置
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {

		j := n - 1
		// 找到 比 i 大的数
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		// 交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 逆序
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}