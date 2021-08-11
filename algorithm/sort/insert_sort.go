package sort

func insertSort(list []int) {
	for i := 0; i < len(list); i++ {
		tmp := list[i]
		for j := i - 1; j >= 0; j-- {
			if tmp > list[j] {
				break
			}
			list[j+1], list[j] = list[j], tmp
		}
	}
}
