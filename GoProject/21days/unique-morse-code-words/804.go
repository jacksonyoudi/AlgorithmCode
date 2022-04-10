package unique_morse_code_words

func uniqueMorseRepresentations(words []string) int {

	wordsMap := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	result := make(map[string]int, 0)

	for _, v := range words {
		bytes := []byte(v)
		tmp := ""
		for _, value := range bytes {
			s := wordsMap[value-'a']
			tmp = tmp + s
		}

		_, ok := result[tmp]
		if ok {
			continue
		} else {
			result[tmp] = 1
		}
	}
	return len(result)

}
