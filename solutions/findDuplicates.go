package solutions

func FindDuplicates(nums []int) []int {
	sort(nums)
	var res []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, nums[i])
		}
	}
	return res
}

// Input: nums = [4,3,2,7,8,2,3,1]
// 7	3	2	4	8	2	3	1
// 3	3	2	4	8	2	7	1
// 2	3	3	4	8	2	7	1
// 3	2	3	4	8	2	7	1
// 3	2	3	4	1	2	7	8
// 1	2	3	4	3	2	7	8
// 1	2	3	4	3	2	7	8
// so element not at correct index is duplicate

func sort(nums []int) {
	index := 0
	for index < len(nums) {
		correctIndex := nums[index] - 1
		if nums[index] != nums[correctIndex] {
			swap(nums, index, correctIndex)
		} else {
			index++
		}
	}
}

func swap(nums []int, index int, correctIndex int) {
	nums[index], nums[correctIndex] = nums[correctIndex], nums[index]
}
