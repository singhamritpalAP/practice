package solutions

func Rotate(matrix [][]int) [][]int {
	transpose(matrix)
	reverseRows(matrix)
	return matrix
}

func transpose(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			swapMatrix(matrix, i, j)
		}
	}
}

// n x n matrix
func swapMatrix(matrix [][]int, i int, j int) {
	if i == j {
		return
	}
	temp := matrix[i][j]
	matrix[i][j] = matrix[j][i]
	matrix[j][i] = temp
}

func reverseRows(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i])/2; j++ {
			swap(matrix[i], j, len(matrix[i])-j-1)
		}
	}
}

//func swap(nums []int, index int, correctIndex int) {
//	nums[index], nums[correctIndex] = nums[correctIndex], nums[index]
//}
