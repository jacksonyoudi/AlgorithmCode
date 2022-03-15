package array

func removeDuplicates(nums []int) int {
	slow, quick := 0, 0

	l := len(nums)
	for i := 0; i < l-1; i++ {
		quick += 1
		if nums[slow] == nums[quick] {
			continue
		} else {
			slow += 1
			if slow == quick {
				continue
			} else {
				nums[slow] = nums[quick]
			}
		}
	}

	return slow+1
}
