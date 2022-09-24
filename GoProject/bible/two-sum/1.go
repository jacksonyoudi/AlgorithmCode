package two_sum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		c := target - v
		if index, ok := m[c]; ok {
			return []int{index, i}
		} else {
			m[v] = i
		}
	}
	return nil
}
