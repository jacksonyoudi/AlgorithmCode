package main

var m = []string{
	"",     // 0
	"",     // 1
	"abc",  // 2
	"def",  // 3
	"ghi",  // 4
	"jkl",  // 5
	"mno",  // 6
	"pqrs", // 7
	"tuv",  // 8
	"wxyz", // 9
}

func letterBackTrace(a string, index int,  box string, anwers []string) {
	ll := len(a)

	if len(box) == ll {
		anwers = append(anwers, box)
	}

	// 第i个人
	stoneIndex :=  int([]byte(a)[index])

	r := m[stoneIndex]

	rb := []byte(r)

	for i := 0; i < len(rb); i++ {
		stone := rb[i]
		box = box + string(stone)
		letterBackTrace(a, index + 1, box,  anwers)

		// 恢复

	}
}