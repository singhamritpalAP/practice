package solutions

func RepeatedAndMissingNumber(A []int) []int {
	sortArray(A)
	var result []int
	for i := 0; i < len(A); i++ {
		if A[i] != i+1 {
			result = append(result, A[i], i+1)
			return result
		}
	}
	return result
}

func sortArray(arr []int) {
	for i := 0; i < len(arr); {
		correctIndex := arr[i] - 1 // 1->0, 2->1
		if arr[i] != i+1 && arr[i] != arr[correctIndex] {
			swap(arr, i, correctIndex)
		} else {
			i++
		}
	}
}

//func swap(nums []int, index int, correctIndex int) {
//	nums[index], nums[correctIndex] = nums[correctIndex], nums[index]
//}

/*
Input:[3 1 2 5 3]
[2 1 3 5 3]
[1 2 3 5 3]
[1 2 3 3 5]

[3 1 3 2 5]
[1 3 3 2 5]
[1 3 3 2 5]
[1 2 3 3 5]
*/
