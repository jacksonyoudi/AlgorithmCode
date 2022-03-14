package slidewindow

import "math"

func minWindow(s string, t string) string {
	need, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0
	// 记录最小覆盖子串的起始索引及长度
	start := 0
	sLen := len(s)
	mlen := math.MaxInt

	for right < sLen {
		// c 是将移入窗口的字符
		c := s[right]
		// 右移窗口
		right++
		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			// 在这里更新最小覆盖子串
			if right-left < mlen {
				start = left
				mlen = right - left
			}

			// d 是将移出窗口的字符
			d := s[left]
			// 左移窗口
			left++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	// 返回最小覆盖子串
	if mlen == math.MaxInt {
		return ""
	} else {
		return s[start : mlen+start]
	}
}
