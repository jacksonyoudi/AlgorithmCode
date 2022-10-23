package _0221023

func helper(nums []int, k int) bool {
	if findGCD(nums) == k {
		return true
	} else {
		return false
	}

}

func findGCD(nums []int) int {
	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return gcd(max, min)
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func subarrayGCD(nums []int, k int) int {
	cnt := 0
	l := len(nums)

	for i := 0; i < l-1; i++ {
		for j := i; j < l; j++ {
			if helper(nums[i:j+1], k) {
				cnt += 1
			}
		}
	}

	return cnt

}
