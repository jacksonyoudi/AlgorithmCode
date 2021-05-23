package array

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, v := range nums {
		j, ok := m[target - v]
		if ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}
