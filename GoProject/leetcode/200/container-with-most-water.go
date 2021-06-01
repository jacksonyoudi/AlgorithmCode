package _200

func maxArea(height []int) int {
	m := 0
	lm := len(height) - 1

	l1, l2 := 0, lm

	for l1 < lm && l2 > 0 && l1 < l2 {
		low := 0
		if height[l1] < height[l2] {
			low = height[l1]
			l1++
		} else {
			low = height[l2]
			l2--
		}

		area := low * (l2 - l1 + 1)
		if area > m {
			m = area
		}
	}
	return m

}
