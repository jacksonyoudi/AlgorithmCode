package reverse_words_in_a_string_iii

func reverseWords(s string) string {
	all := []byte(s)

	tmp := make([]byte, 0)
	result := ""
	for i := 0; i < len(all); i++ {
		if all[i] == ' ' {
			s := string(tmp)
			s2 := helper(s)
			result = result + s2 + " "
			tmp = make([]byte, 0)
		} else {
			tmp = append(tmp, all[i])
		}
	}
	s1 := string(tmp)
	s2 := helper(s1)
	result = result + s2

	return result
}

func helper(str string) string {
	s := []byte(str)
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
	return string(s)
}
