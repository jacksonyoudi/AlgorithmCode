package bit

func maxProduct(words []string) int {
	if words == nil || len(words) < 2 {
		return 0
	}
	length := len(words)
	masks := make([]int, length)
	for i, word := range words[:] {
		for _, c := range word {
			masks[i] |= 1 << uint(c - 'a')
		}
	}


	max := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			// 完全没有相同的单词
			if (masks[i] & masks[j]) == 0 {
				prod := len(words[i]) * len(words[j])
				if prod > max {
					max = prod
				}
			}
		}
	}
	return max
}