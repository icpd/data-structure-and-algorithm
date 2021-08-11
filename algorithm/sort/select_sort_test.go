package sort

import "testing"

func Test_selectSort(t *testing.T) {
	s1 := []int{3, 343, 5436, 5, 23}
	s2 := []int{34, 0, 3, 4, 1}

	t.Run("s1", func(t *testing.T) {
		selectSort(s1)
		t.Log(s1)
	})

	t.Run("s2", func(t *testing.T) {
		selectSort(s2)
		t.Log(s2)
	})

}
