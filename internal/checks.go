package internal

func checkNumber(sudoku [9][9]int8, n, row, column int8) bool {
	return checkBox(sudoku, n, row, column) && checkLine(sudoku, n, row, column) && checkColumn(sudoku, n, row, column)
}

func checkBox(sudoku [9][9]int8, n, row, column int8) bool {
	minRow, maxRow := getBounds(row)
	minColumn, maxColumn := getBounds(column)

	var i int8
	for i = minRow; i <= maxRow; i++ {
		var j int8
		for j = minColumn; j <= maxColumn; j++ {
			if i == row && j == column {
				continue
			}

			if sudoku[i][j] == n {
				return false
			}
		}
	}

	return true
}

func checkLine(sudoku [9][9]int8, n, row, column int8) bool {
	var i int8
	for i = 0; i < 9; i++ {
		if i == row {
			continue
		}

		if sudoku[i][column] == n {
			return false
		}
	}

	return true
}

func checkColumn(sudoku [9][9]int8, n, row, column int8) bool {
	var i int8
	for i = 0; i < 9; i++ {
		if i == column {
			continue
		}

		if sudoku[row][i] == n {
			return false
		}
	}

	return true
}

func getBounds(axis int8) (int8, int8) {
	if axis < 3 {
		return 0, 2
	} else if axis < 6 {
		return 3, 5
	} else {
		return 6, 8
	}
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
