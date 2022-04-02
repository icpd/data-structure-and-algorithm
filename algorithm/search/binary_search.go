package search

func BinarySearch(target int, sli []int) int {
	if len(sli) <= 0 {
		return -1
	}
	right := len(sli) - 1
	left := 0
	for left <= right {
		mid := (left + right) / 2
		if sli[mid] == target {
			return mid
		}
		if target < sli[mid] {
			right = mid - 1
		}
		if target > sli[mid] {
			left = mid + 1
		}
	}
	return -1
}

// 查找第一个值等于给定值的元素
func BSearch1(target int, sli []int) int {
	if len(sli) <= 0 {
		return -1
	}
	high := len(sli) - 1
	low := 0

	for low <= high {
		mid := low + ((high - low) >> 1)

		if sli[mid] > target {
			high = mid - 1
		}

		if sli[mid] < target {
			low = mid + 1
		}

		if sli[mid] == target {
			if mid == 0 || sli[mid-1] < target {
				return mid
			}

			// 如果上一个值不是目标值，就继续查找
			high = mid - 1
		}
	}

	return -1
}

// 查找最后一个值等于给定的元素
func BSearch2(target int, sli []int) int {
	if len(sli) <= 0 {
		return -1
	}
	high := len(sli) - 1
	low := 0

	for low <= high {
		mid := low + ((high - low) >> 1)

		if sli[mid] > target {
			high = mid - 1
		}

		if sli[mid] < target {
			low = mid + 1
		}

		if sli[mid] == target {
			if mid == len(sli)-1 || sli[mid+1] > target {
				return mid
			}

			low = mid + 1
		}
	}

	return -1
}

// 查找第一个大于等于给定值的元素
func BSearch3(target int, sli []int) int {
	if len(sli) <= 0 {
		return -1
	}
	high := len(sli) - 1
	low := 0

	for low <= high {
		mid := low + ((high - low) >> 1)

		if sli[mid] >= target {
			if mid == 0 || sli[mid-1] < target {
				return mid
			}

			high = mid - 1
		}

		if sli[mid] < target {
			low = mid + 1
		}
	}

	return -1
}

// 查找最后一个小于等于给定元素的值
func BSearch4(target int, sli []int) int {
	if len(sli) <= 0 {
		return -1
	}
	high := len(sli) - 1
	low := 0

	for low <= high {
		mid := low + ((high - low) >> 1)

		if sli[mid] > target {
			high = mid - 1
		}

		if sli[mid] <= target {
			if mid == len(sli)-1 || sli[mid+1] > target {
				return mid
			}

			low = mid + 1
		}
	}

	return -1
}
