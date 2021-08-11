package sort

import "testing"

func Test_insertSort(t *testing.T) {
	s1 := []int{123, 343, 34, 454, 2, 1}
	t.Run("test1", func(t *testing.T) {
		insertSort(s1)
		t.Log(s1)
	})
}
