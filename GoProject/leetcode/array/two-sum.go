package array

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	l := len(nums)

	for i := 0; i < l; i++ {
		t := target - nums[i]

		v, ok := m[t]
		if ok {
			return []int{v, i}
		} else {
			m[nums[i]] = i
		}
	}
	return nil

}
