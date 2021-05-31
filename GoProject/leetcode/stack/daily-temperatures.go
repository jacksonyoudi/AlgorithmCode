package stack

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			// 出栈
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}
