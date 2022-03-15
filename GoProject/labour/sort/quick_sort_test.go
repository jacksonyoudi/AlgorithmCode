package sort

import "testing"



func Test_quick_sort(t *testing.T) {
	type args struct {
		nums  []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"one",
			args{
				[]int{10, 7, 8, 9, 1, 5},
				0,
				6,
			},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
