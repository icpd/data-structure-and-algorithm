package search

import "testing"

func TestBinarySearch(t *testing.T) {
	sli := []int{0, 1, 2, 3, 5, 6, 7, 8, 9}
	t.Log(BinarySearch(7, sli))
}
