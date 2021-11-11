package main


func chalkReplacer(chalk []int, k int) int {

	s := k

	if k == 0 {
		return 0
	}

	m := -1
	for s >= 0 {

		for i,v := range chalk {
			s = s -v
			if s < 0 {
				m = i
				return m
			}
		}
	}

	return m

}
