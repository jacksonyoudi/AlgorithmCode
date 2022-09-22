package check_array_formation_through_concatenation

func canFormArray(arr []int, pieces [][]int) bool {

	cnt := 0
	lc := len(arr)
	for _, vs := range pieces {
		l := len(vs)
		start := vs[0]
		index := 0
		first := false
		for _, v := range arr {
			if v != start {
				if first == false {
					continue
				} else {
					break
				}
			}
			first = true
			cnt = cnt + 1
			index = index + 1
			if index == l {
				break
			}
			start = vs[index]
		}
	}

	return lc == cnt

}
