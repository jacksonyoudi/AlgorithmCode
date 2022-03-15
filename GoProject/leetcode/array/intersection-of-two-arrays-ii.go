package array

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for _, v := range nums1 {
		c, ok := m1[v]
		if !ok {
			m1[v] = 1
		} else {
			m1[v] = c + 1
		}
	}

	for _, v := range nums2 {
		c, ok := m2[v]
		if !ok {
			m2[v] = 1
		} else {
			m2[v] = 1 + c
		}
	}

	result := []int{}

	for k, c1 := range m1 {
		c2, ok := m2[k]
		if ok {
			mx := 0
			if c1 <= c2 {
				mx = c1
			} else {
				mx = c2
			}

			for i := 0; i < mx; i++ {
				result = append(result, k)
			}
		}
	}
	return result
}
