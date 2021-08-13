package sort

import (
	"testing"
)

func Test_quickSort(t *testing.T) {
	sli := []int{1, 3, 5, 8, 10, 89, 9, 0, 12}
	t.Log(quickSort(sli))
}
