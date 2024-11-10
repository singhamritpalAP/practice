package solutions

import "fmt"

// LC - 75
func SortColors(nums []int, approach int) {
	if approach == 1 {
		approach1(nums)
	} else if approach == 2 {
		approach2(nums)
	}
}

func approach1(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		j = swapIfNeeded(nums, i, j, 0)
	}
	if j < len(nums) && nums[j] == 0 {
		j++
	}
	for i := j; i < len(nums); i++ {
		j = swapIfNeeded(nums, i, j, 1)
	}
	if j < len(nums) && nums[j] == 1 {
		j++
	}
	for i := j; i < len(nums); i++ {
		j = swapIfNeeded(nums, i, j, 2)
	}
	fmt.Println(nums)
}

func approach2(nums []int) {
	zeros, ones, twos := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeros++
		} else if nums[i] == 1 {
			ones++
		} else if nums[i] == 2 {
			twos++
		}
	}
	i := 0
	for zeros > 0 {
		nums[i] = 0
		zeros--
		i++
	}
	for ones > 0 {
		nums[i] = 1
		ones--
		i++
	}
	for twos > 0 {
		nums[i] = 2
		twos--
		i++
	}
	fmt.Println(nums)
}

// func swap(nums []int, index int, correctIndex int) {
//	nums[index], nums[correctIndex] = nums[correctIndex], nums[index]
//}

func swapIfNeeded(nums []int, i int, j int, value int) int {
	if nums[i] == value && nums[j] != value {
		swap(nums, i, j)
		j++
	} else if nums[i] == value {
		j++
	}
	return j
}

/*
[2,0,2,1,1,0]
 j
 i
if 0; & nums[j] != 0 -> swap followed by i,j++
else if 0; j++
if !=0 i++
same for 1, 2
[2,0,2,1,1,0]
[2,0,2,1,1,0]
[0,2,2,1,1,0]
[0,0,2,1,1,2]
[0,0,1,2,1,2]
[0,0,1,1,2,2]

*/
