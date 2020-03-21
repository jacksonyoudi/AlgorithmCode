package leetcode

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	maxTmp := 0
	max := 0
	for ; start < end; {
		if height[start] < height[end] {
			maxTmp = height[start] * (end - start)
			start++
		} else {
			maxTmp = height[end] * (end - start)
			end--
		}

		if maxTmp > max {
			max = maxTmp
		}
	}

	return max
}
