package longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	l := len(nums)
	dp := make([]int, l)

	for i := 0; i < l; i++ {
		dp[i] = 1
	}

	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	res := 0

	for i := 0; i < l; i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
