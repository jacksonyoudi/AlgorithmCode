package array

func replaceSpace(s string) string {
	bs := []byte(s)
	result := make([]byte, 0, len(bs))
	for _, b := range bs {
		if b == ' ' {
			m := []byte("%20")
			result = append(result, m...)
		} else {
			m := []byte{b}
			result = append(result, m...)
		}
	}
	return string(result)
}
