package main

// https://leetcode.com/problems/set-matrix-zeroes

func SetMatrixZeroes(matrix [][]int) {
	setFirstCol, setFirstRow := false, false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 { // j = 0
			setFirstCol = true
		}
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 || j == 0 {
				if matrix[0][j] == 0 { // i = 0
					setFirstRow = true
				}
			} else {
				if matrix[i][j] == 0 {
					matrix[0][j] = 0
					matrix[i][0] = 0
				}
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if setFirstCol {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

	if setFirstRow {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
}
