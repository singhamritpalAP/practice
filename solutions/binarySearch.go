package solutions

func Search(arr []int, target int, low int, high int) int {
	for low <= high {
		mid := low + ((high - low) / 2)
		if arr[mid] == target {
			return mid + 1
		}
		if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
