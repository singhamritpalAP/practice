package solutions

func Generate(numRows int, approach int) [][]int {
	if approach == 1 {
		return pascalApproach1(numRows)
	} else {
		return pascalApproach2(numRows)
	}
}

func pascalApproach1(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		result[i-1] = make([]int, i)
		for j := 1; j <= i; j++ {
			result[i-1][j-1] = nCr(i-1, j-1)
		}
	}
	return result
}
func nCr(n, r int) int {
	result := 1
	for i := 0; i < r; i++ {
		result = (result * (n - i)) / (i + 1)
	}
	return result
}

func pascalApproach2(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		result[i-1] = generateRow(i)
	}
	return result
}

func generateRow(row int) []int {
	ans := 1
	result := make([]int, row)
	result[0] = ans
	for i := 1; i < row; i++ {
		ans = (ans * (row - i)) / i
		result[i] = ans
	}
	return result
}
