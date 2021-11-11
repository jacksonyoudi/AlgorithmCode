package dp

import "math"

func coinChange(coins []int, amount int) int {

	INF := math.MaxInt32
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = INF
	}

	// dp初始条件
	dp[0] = 0

	for i := 0; i < amount; i++ {
		for _, y := range coins {
			if y <= amount && i+y < amount && i+y >= 0 {
				if dp[i+y] >= dp[i]+1 {
					dp[i+y] = dp[i] + 1
				} else {
					dp[i+y] = dp[i+y]
				}
			}
		}
	}
	if dp[amount] >= INF {
		return -1
	} else {
		return dp[amount]
	}

}
