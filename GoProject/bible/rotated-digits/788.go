package rotated_digits

var goodNums = map[int]int{
	2: 5,
	5: 2,
	6: 9,
	9: 6,
	0: 0,
	8: 8,
	1: 1,
}

func helper(n int) bool {
	high := n / 10
	low := n % 10

	r := 0
	i := 0
	for high > 0 || low > 0 {
		if v, ok := goodNums[low]; ok {

			for j := 0; j < i; j++ {
				v = v * 10
			}

			r = r + v
			h := high
			high = h / 10
			low = h % 10
			i++
			continue
		} else {
			return false
		}
	}
	if r != n {
		return true
	} else {
		return false
	}
}

func rotatedDigits(n int) int {
	cnt := 0
	for i := 1; i <= n; i++ {
		if helper(i) {
			cnt++
		}
	}
	return cnt
}
