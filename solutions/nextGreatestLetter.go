package solutions

func NextGreatestLetter(letters []byte, target byte) byte {
	low := 0
	high := len(letters) - 1
	if target > letters[high] {
		return letters[0]
	}
	for low <= high {
		mid := low + (high-low)/2
		if letters[mid] == target && mid == len(letters)-1 {
			return letters[0]
		} else if letters[mid] == target && letters[mid+1] != target {
			return letters[mid+1]
		}
		if letters[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return letters[low]
}
