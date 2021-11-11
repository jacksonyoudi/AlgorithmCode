package main

type PermuteData struct {
	d [][]int
}

func permuteDfs(nums []int, index int, box []int, anwers *PermuteData) {
	if index == len(box) {
		anwers.d = append(anwers.d, box)
		return
	}

	for i, v := range nums {
		m := []int{}
		m = append(m, box...)
		m = append(m, v)

		newNums := []int{}
		newNums = append(newNums, nums[0:i]...)
		newNums = append(newNums, nums[i+1:]...)

		permuteDfs(newNums, index, m, anwers)

	}
}

func permute(nums []int) [][]int {
	d := &PermuteData{d: make([][]int, 0)}
	tmp := make([]int, 0)

	permuteDfs(nums, len(nums), tmp, d)
	return d.d
}
