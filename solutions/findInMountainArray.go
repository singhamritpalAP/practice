package solutions

// MountainArray LC: 1095/**
type MountainArray struct {
}

// These two are dummy functions; code executed in LC
func (this *MountainArray) get(index int) int {
	return 0
}
func (this *MountainArray) length() int {
	return 10
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	length := mountainArr.length()
	low, high := 0, length-1
	peakIndex := findPeakIndex(mountainArr, low, high, length)
	if peakIndex == -1 {
		return -1
	}

	if resultIndex := searchInAnyOrder(mountainArr, 0, peakIndex, true, target); resultIndex != -1 {
		return resultIndex
	} else {
		return searchInAnyOrder(mountainArr, peakIndex, length-1, false, target)
	}
}

func searchInAnyOrder(arr *MountainArray, low int, high int, isAscending bool, target int) int {
	for low <= high {
		mid := low + (high-low)/2
		if arr.get(mid) == target {
			return mid
		}
		if isAscending {
			if arr.get(mid) < target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if arr.get(mid) > target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func findPeakIndex(arr *MountainArray, low int, high int, length int) int {
	for low < high {
		mid := low + (high-low)/2
		if mid+1 < length && arr.get(mid) > arr.get(mid+1) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
