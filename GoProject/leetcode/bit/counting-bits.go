package bit

func countBits(num int) []int {
	dp := make([]int, num+1)

	for i := 1; i < num; i++ {
		if i&1 == 1 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i>>1]
		}
	}
	return dp
}
