package array

func plusOne(digits []int) []int {
	l := len(digits)
	tmp := 1
	for i := 0; i < l; i++ {
		if (digits[l-1-i] + tmp) == 10 {
			digits[l-1-i] = 0
			tmp = 1
		} else {
			digits[l-1-i] = digits[l-1-i] + tmp
			tmp = 0
		}
	}
	if tmp == 1 {
		result := []int{1}
		result = append(result, digits...)
		return result
	}
	return digits


}
