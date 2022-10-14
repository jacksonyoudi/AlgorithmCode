package check_if_binary_string_has_at_most_one_segment_of_ones

func checkOnesSegment(s string) bool {
	var pre byte = '1'
	for _, v := range []byte(s) {
		if v == pre || (v == '0' && pre == '1') {
			pre = v
			continue
		} else {
			return false
		}
	}
	return true
}
