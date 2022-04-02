package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	sli := []int{0, 1, 2, 3, 5, 6, 7, 8, 9}
	t.Log(BinarySearch(7, sli))
}

func TestBSearch1(t *testing.T) {
	sli := []int{1, 2, 3, 4, 5, 5, 5, 6, 7, 8}
	assert.Equal(t, 4, BSearch1(5, sli))
	assert.Equal(t, 7, BSearch1(6, sli))
	assert.Equal(t, -1, BSearch1(10, sli))
	assert.Equal(t, 0, BSearch1(1, sli))
}

func TestBSearch2(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 5, 5, 6, 7, 8}
	type args struct {
		target int
		sli    []int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{1, s},
			want: 0,
		},
		{
			name: "t2",
			args: args{5, s},
			want: 6,
		},
		{
			name: "t3",
			args: args{8, s},
			want: 9,
		},
		{
			name: "t4",
			args: args{111, s},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, BSearch2(tt.args.target, tt.args.sli), "BSearch2(%v, %v)", tt.args.target, tt.args.sli)
		})
	}
}

func TestBSearch3(t *testing.T) {
	sli := []int{1, 2, 3, 4, 5, 5, 5, 6, 8, 9}
	assert.Equal(t, 4, BSearch3(5, sli))
	assert.Equal(t, 8, BSearch3(7, sli))
	assert.Equal(t, -1, BSearch3(10, sli))
	assert.Equal(t, 0, BSearch3(1, sli))
}

func TestBSearch4(t *testing.T) {
	sli := []int{1, 2, 3, 4, 5, 5, 5, 6, 8, 9}
	assert.Equal(t, 6, BSearch4(5, sli))
	assert.Equal(t, 7, BSearch4(7, sli))
	assert.Equal(t, 9, BSearch4(10, sli))
	assert.Equal(t, 0, BSearch4(1, sli))
}
