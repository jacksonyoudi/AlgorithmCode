package str

func reverseString(s []byte) {
	l := len(s)

	for i := 0; i < l/2; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
}
