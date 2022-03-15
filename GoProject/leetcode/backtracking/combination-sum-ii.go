package main

func combinationSumDfs2(candidates []int, index int, box []int, anwers *combinationSumData, target int) {
	ll := len(candidates)
	if index >= ll || sum(box) > target {
		return
	}

	if sum(box) == target {
		anwers.d = append(anwers.d, box)
		return
	}

	for i := index; i < ll; i++ {
		m := []int{}
		m = append(m, box...)
		m = append(m, candidates[i])
		combinationSumDfs2(candidates, i+1, m, anwers, target)
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	anwers := &combinationSumData{
		d: [][]int{},
	}

	combinationSumDfs2(candidates, 0, []int{}, anwers, target)
	return anwers.d
}
