package str

func firstUniqChar(s string) byte {

	m := make(map[byte]int)
	bytes := []byte(s)

	for i, b := range bytes {
		_, ok := m[b]
		if ok {
		} else {
			m[b] = i
		}
	}

	if len(m) == 0 {
		return ' '
	}

	min := len(bytes)

	for _, v := range m {
		if min > v {
			min = v
		}
	}
	return bytes[min]
}
