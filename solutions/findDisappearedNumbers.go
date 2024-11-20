package solutions

func FindDisappearedNumbers(nums []int) []int {
	sortArr(nums)
	var res []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}
