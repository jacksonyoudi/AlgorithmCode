package sort

import (
	"reflect"
	"testing"
)

func Test_merge_sort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"one",
			args{
				[]int{10, 7, 8, 9, 1, 5},
			},
			[]int{1, 5, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge_sort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge_sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
