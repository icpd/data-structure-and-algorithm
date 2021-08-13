package sort

func quickSort(sli []int) []int {
	//先判断是否需要继续进行
	length := len(sli)
	if length <= 1 {
		return sli
	}
	//选择第一个元素作为基准
	baseNum := sli[0]
	//遍历除了标尺外的所有元素，按照大小关系放入左右两个切片内
	//初始化左右两个切片
	var leftSli []int  //小于基准的
	var rightSli []int //大于基准的
	for i := 1; i < length; i++ {
		if baseNum > sli[i] {
			//放入左边切片
			leftSli = append(leftSli, sli[i])
		} else {
			//放入右边切片
			rightSli = append(rightSli, sli[i])
		}
	}

	//再分别对左边和右边的切片进行相同的排序处理方式递归调用这个函数
	leftSli = quickSort(leftSli)
	rightSli = quickSort(rightSli)

	//合并
	leftSli = append(leftSli, baseNum)
	return append(leftSli, rightSli...)
}

// ==================== 其他版本 ======================
func quickSortV2(sli []int) {
	recursionSort(sli, 0, len(sli)-1)
}

func recursionSort(sli []int, left int, right int) {
	if left < right {
		pivot := partition(sli, left, right)
		recursionSort(sli, left, pivot-1)
		recursionSort(sli, pivot+1, right)
	}
}

func partition(sli []int, left int, right int) int {
	for left < right {
		for left < right && sli[left] <= sli[right] {
			right--
		}
		if left < right {
			sli[left], sli[right] = sli[right], sli[left]
			left++
		}

		for left < right && sli[left] <= sli[right] {
			left++
		}
		if left < right {
			sli[left], sli[right] = sli[right], sli[left]
			right--
		}
	}

	return left
}
