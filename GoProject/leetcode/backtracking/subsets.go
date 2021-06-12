package main

type subsetsData struct {
	d [][]int
}

func subseetsBackTrace(nums []int, index int, box []int, anwers *subsetsData) {
	ll := len(nums)
	// 第一次可以有一个空
	anwers.d = append(anwers.d, box)
	if index >= ll {
		return
	}

	for i := index; i < ll; i++ {
		m := []int{}
		m = append(m, box...)
		m = append(m, nums[i])
		subseetsBackTrace(nums, i+1, m, anwers)
	}

}

func subsets(nums []int) [][]int {
	anwers := &subsetsData{d: [][]int{}}
	box := []int{}

	subseetsBackTrace(nums, 0, box, anwers)
	return anwers.d
}
