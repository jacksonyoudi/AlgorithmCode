package bit


func hammingDistance(x int, y int) int {
	diff := x ^ y
	ans := 0

	for diff > 0 {
		ans += diff & 1
		diff >>= 1
	}
	return ans
}
