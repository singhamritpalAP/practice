package solutions

import (
	"sort"
)

func Merge(intervals [][]int) [][]int {
	result := make([][]int, 1)
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result[0] = intervals[0]
	index := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[index][1] >= intervals[i][0] {
			intervals[index][1] = max(intervals[index][1], intervals[i][1])
		} else {
			index++
			intervals[index] = intervals[i]
		}
	}
	return intervals[:index+1]
}

func max(ele1 int, ele2 int) int {
	if ele1 > ele2 {
		return ele1
	} else {
		return ele2
	}
}
