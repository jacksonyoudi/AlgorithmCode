package main

func reverse_string(s []byte) {
	if len(s) <= 1 {
		return
	}


	s = append(s[1:], s[0])
	reverse_string(s)

}
