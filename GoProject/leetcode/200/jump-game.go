package _200

func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	maxJump := 0
	for i, v := range nums {
		if i > maxJump {
			return false
		}

		if maxJump < i+v {
			maxJump = i + v
		}
	}
	return true
}
