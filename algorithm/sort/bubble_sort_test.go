package sort

import (
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{[]int{1, 9, 3, 10, 7}}},
		{name: "2", args: args{[]int{127, 34, 0, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubbleSort(tt.args.a)
			t.Log(tt.args.a)
		})
	}
}
