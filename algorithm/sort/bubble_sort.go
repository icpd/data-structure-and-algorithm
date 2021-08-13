package sort

func bubbleSort(sli []int) []int {
	length := len(sli)
	//该层循环控制 需要冒泡的轮数
	for i := 1; i < length; i++ {
		//该层循环用来控制每轮 冒出一个数 需要比较的次数
		for j := 0; j < length-1; j++ {
			if sli[i] < sli[j] {
				sli[i], sli[j] = sli[j], sli[i]
			}
		}
	}
	return sli
}
