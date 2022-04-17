package valid_parentheses

func isValid(s string) bool {
	l := len(s)
	if l%2 == 1 {
		return false
	}
	// 数组模拟栈

	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]byte, 0)

	for _, v := range []byte(s) {
		m, ok := pairs[v]
		if ok {
			if len(stack) == 0 || m != stack[len(stack)-1] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, v)
		}
	}

	return len(stack) == 0

}
