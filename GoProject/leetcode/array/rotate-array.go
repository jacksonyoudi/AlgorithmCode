package array

func rotate(nums []int, k int) {

	// 移动一个

	l := len(nums)

	for j:=0;j<k;j++ {
		tmp := nums[l-1]

		for i := 0; i < l-1; i++ {
			nums[l-1-i] = nums[l-2-i]
		}
		nums[0] = tmp
	}

}
