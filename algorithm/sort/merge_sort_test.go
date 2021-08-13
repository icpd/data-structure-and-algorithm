package sort

import (
	"testing"
)

func Test_mergeSort(t *testing.T) {
	s := []int{6, 7, 8, 10, 4, 6, 99}
	res := mergeSort(s)
	t.Log(res)
}
