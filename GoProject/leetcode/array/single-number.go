package array

func singleNumber(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		c, ok := m[v]
		if ok {
			m[v] = c + 1
		} else {
			m[v] = 1
		}
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}
