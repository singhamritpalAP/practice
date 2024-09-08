package solutions

func SearchRange(nums []int, target int) []int {
	//if len(nums) == 0 || (len(nums) == 1 && nums[0] != target) {
	//	return []int{-1, -1}
	//}
	//if len(nums) == 1 && nums[0] == target {
	//	return []int{0, 0}
	//}

	res := []int{-1, -1}
	res[0] = search(nums, target, true)
	res[1] = search(nums, target, false)
	return res
}

func search(nums []int, target int, isLeftSearch bool) int {
	low, high, index := 0, len(nums)-1, -1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target && isLeftSearch {
			index = mid
			high = mid - 1
		}
		if nums[mid] == target && !isLeftSearch {
			index = mid
			low = mid + 1
		}

		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		}

	}
	return index
}
