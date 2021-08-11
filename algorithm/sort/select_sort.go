package sort

func selectSort(list []int) {
	length := len(list)
	for i := 0; i < length; i++ {
		maxIndex := 0
		//寻找最大的一个数，保存索引值
		for j := 1; j < length-i; j++ {
			if list[j] > list[maxIndex] {
				maxIndex = j
			}
		}
		list[length-i-1], list[maxIndex] = list[maxIndex], list[length-i-1]
	}
}
