package algorithm

func BinarySearch(target int, sli []int) int {
	if len(sli) < 0 {
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
