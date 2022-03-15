package _200

var m = map[string][]string{
	"2": []string{
		"a",
		"b",
		"c",
	},
	"3": []string{
		"d",
		"e",
		"f",
	},
	"4": []string{
		"g",
		"h",
		"i",
	},
	"5": []string{
		"j",
		"k",
		"l",
	},
	"6": []string{
		"m",
		"n",
		"o",
	},
	"7": []string{
		"p",
		"q",
		"r",
		"s",
	},
	"8": []string{
		"t",
		"u",
		"v",
	},
	"9": []string{
		"w",
		"x",
		"y",
		"z",
	},
}

func letterCombinationsHelper(data []string, digits string) []string {
	if len(digits) == 0 {
		return data
	}

	a := digits[0]
	words, _ := m[string(a)]
	result := make([]string, 0)
	for _, v := range words {
		if len(data) > 0 {
			for _, d := range data {
				result = append(result, d+v)
			}
		} else {
			result = append(result, v)
		}
	}
	return letterCombinationsHelper(result, digits[1:])
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := make([]string, 0)
	return  letterCombinationsHelper(result, digits)
}
