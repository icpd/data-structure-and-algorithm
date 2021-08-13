package sort

func selectSort(sli []int) []int {
	//双重循环完成，外层控制轮数，内层控制比较次数
	length := len(sli)
	for i := 0; i < length-1; i++ {
		//先假设最小的值的位置
		k := i
		for j := i + 1; j < length; j++ {
			//sli[k] 是当前已知的最小值
			if sli[k] > sli[j] {
				//比较，发现更小的,记录下最小值的位置；并且在下次比较时采用已知的最小值进行比较。
				k = j
			}
		}
		//已经确定了当前的最小值的位置，保存到 k 中。如果发现最小值的位置与当前假设的位置 i 不同，则位置互换即可。
		if k != i {
			sli[k], sli[i] = sli[i], sli[k]
		}
	}
	return sli
}
