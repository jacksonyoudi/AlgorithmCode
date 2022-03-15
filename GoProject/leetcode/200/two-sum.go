package _200

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		t := target - v
		j, ok := m[t]
		if ok {
			return []int{i, j}
		} else {
			m[v] = i
		}
	}

	return []int{0, 0}

}
