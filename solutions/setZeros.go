package solutions

import "math"

// LC - 73
func setZeroes(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		isRowSet := false
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				// set entire row to -1
				setRow(matrix, i, isRowSet)
				isRowSet = true
				// set entire col to -1
				setCol(matrix, j)
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == math.MinInt {
				matrix[i][j] = 0
			}
		}
	}
}

func setCol(matrix [][]int, j int) {
	for i := 0; i < len(matrix); i++ {
		if matrix[i][j] != 0 {
			matrix[i][j] = math.MinInt
		}
	}
}

func setRow(matrix [][]int, i int, isRowSet bool) {
	// set row only if not set earlier
	if !isRowSet {
		for col := 0; col < len(matrix[i]); col++ {
			if matrix[i][col] != 0 { // so that if any other element in row is 0, that entire col can be set to -1.
				matrix[i][col] = math.MinInt
			}
		}
	}
}
