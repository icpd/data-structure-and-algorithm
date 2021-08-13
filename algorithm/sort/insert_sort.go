package sort

func insertSort(sli []int) []int {
	length := len(sli)
	for i := 0; i < length; i++ {
		tmp := sli[i]
		//内层循环控制，比较并插入
		for j := i - 1; j >= 0; j-- {
			if tmp < sli[j] {
				//发现插入的元素要小，交换位置，将后边的元素与前面的元素互换
				sli[j+1], sli[j] = sli[j], tmp
			} else {
				//如果碰到不需要移动的元素，则前面的就不需要再次比较了。
				break
			}
		}
	}
	return sli
}
