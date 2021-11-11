package main

type combinationSumData struct {
	d [][]int
}

func sum(d []int) int {
	s := 0
	for _, v := range d {
		s += v
	}
	return s
}

func combinationSumDfs(candidates []int, index int, box []int, anwers *combinationSumData, target int) {
	ll := len(candidates)
	if index >= ll {

	}



}

func combinationSum(candidates []int, target int) [][]int {
	return nil
}
