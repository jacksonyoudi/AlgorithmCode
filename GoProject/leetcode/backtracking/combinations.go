package main

type combineData struct {
	d [][]int
}

func combineDfs(n int, index int, box []int, anwers *combineData, k int) {
	if len(box) == k {
		anwers.d = append(anwers.d, box)
		return
	}
	for i := index; i <= n; i++ {
		m := []int{}
		for _, v := range box {
			m = append(m, v)
		}

		m = append(m, i)

		combineDfs(n, i+1, m, anwers, k)
	}
}

func combine(n int, k int) [][]int {
	anwers := &combineData{
		d: [][]int{},
	}

	combineDfs(n, 1, []int{}, anwers, k)
	return anwers.d
}
