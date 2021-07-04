package stack

func isValid(s string) bool {
	n := len(s)

	//
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []byte{}

	for _, b := range []byte(s) {
		i, ok := pairs[b]
		if ok {
			if len(stack) == 0 || stack[len(stack)-1] != i {
				return false
			} else {
				// 从后出
				stack = stack[:len(stack)-1]
			}

		} else {
			stack = append(stack, b)
		}
	}

	return len(stack) == 0

}
