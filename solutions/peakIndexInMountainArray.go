package solutions

func PeakIndexInMountainArray(arr []int) int {
	low, high := 0, len(arr)-1
	for low < high {
		mid := low + (high-low)/2
		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] { // given array will be mountain
			return mid
		}
		if arr[mid] < arr[mid+1] {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return low
}
