package main

func helper(s byte) []byte {
	if s >= 'a' && s <= 'z' {
		return []byte{s, s - 32}
	}

	if s >= 'A' && s <= 'Z' {
		return []byte{s, s + 32}
	}

	return []byte{s}

}

type letterCaseData struct {
	d []string
}

func letterCasePermutationBackTrace(s string, index int, box string, anwers *letterCaseData) {
	ll := len(s)

	//  条件满足，放到anwers 中
	if len(box) == ll {
		anwers.d = append(anwers.d, box)
		return
	}

	people := s[index]

	choices := helper(people)

	for _, v := range choices {
		box = box + string(v)
		letterCasePermutationBackTrace(s, index+1, box, anwers)
		box = string([]byte(box)[0 : len(box)-1])
	}
}

func letterCasePermutation(s string) []string {
	anwers := &letterCaseData{
		d:  []string{},
	}
	letterCasePermutationBackTrace(s, 0, "", anwers)
	return anwers.d
}
