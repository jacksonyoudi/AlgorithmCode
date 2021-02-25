package array

func moveZeroes(nums []int) {
	slow, quick := 0, 1

	l := len(nums) - 1

	for i := 0; i < l; i++ {
		switch {
		case nums[slow] == 0 && nums[quick] != 0:
			nums[slow], nums[quick] = nums[quick], nums[slow]
			slow += 1
			quick += 1
		case nums[slow] != 0 && nums[quick] == 0:
			slow += 1
			quick += 1
		case nums[slow] == 0 && nums[quick] == 0:
			quick += 1
		case nums[slow] != 0 && nums[quick] != 0:
			slow += 1
			if slow == quick {
				quick += 1
			}
		}
	}

}
