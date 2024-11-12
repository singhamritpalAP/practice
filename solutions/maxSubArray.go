package solutions

import (
	"fmt"
	"math"
)

func MaxSubArraySum(nums []int, approach int) int {
	if len(nums) == 0 {
		return 0
	}
	if approach == 1 {
		return maxSubArraySumNxN(nums)
	} else {
		return maxSubArraySumN(nums)
	}
}

func maxSubArraySumNxN(nums []int) int {
	maxSum := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > maxSum {
				maxSum = sum
				// if question is to return subArr also store indexes(i,j)
			}
		}
	}
	return maxSum
}

func maxSubArraySumN(nums []int) int {
	sum, maxSum, startIndex, currentStartIndex, endIndex := 0, math.MinInt, 0, 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > maxSum {
			maxSum = sum
			startIndex = currentStartIndex
			endIndex = i + 1
			// to return subArr store start and end index
		} else if sum < 0 {
			sum = 0
			currentStartIndex = i + 1
			// to return subArr store i + 1 as start index
		}
	}
	subArr := nums[startIndex:endIndex]
	fmt.Println("Sub array", subArr)
	return maxSum
}
