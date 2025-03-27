package internal

func checkNumber(sudoku [9][9]int8, num, row, column int8) bool {
	for i := int8(0); i < 9; i++ {
		if i == row || i == column {
			continue
		}

		if sudoku[row][i] == num || sudoku[i][column] == num {
			return false
		}
	}

	startRow := row / 3 * 3
	startColumn := column / 3 * 3
	for i := int8(0); i < 3; i++ {
		for j := int8(0); j < 3; j++ {
			if startRow+i == row && startColumn+j == column {
				continue
			}

			if sudoku[startRow+i][startColumn+j] == num {
				return false
			}
		}
	}

	return true
}

func VerifySudoku(sudoku [9][9]int8) bool {
	var i int8
	for i = 0; i < 9; i++ {
		var k int8
		for k = 0; k < 9; k++ {
			number := sudoku[i][k]

			if number == 0 {
				continue
			}

			if !checkNumber(sudoku, number, i, k) {
				return false
			}
		}
	}

	return true
}
