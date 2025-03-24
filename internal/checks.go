package internal

func checkNumber(sudoku [9][9]int, n, row, column int) bool {
	return checkBox(sudoku, n, row, column) && checkLine(sudoku, n, row, column) && checkColumn(sudoku, n, row, column)
}

func checkBox(sudoku [9][9]int, n, row, column int) bool {
	minRow, maxRow := getBounds(row)
	minColumn, maxColumn := getBounds(column)

	for i := minRow; i <= maxRow; i++ {
		for j := minColumn; j <= maxColumn; j++ {
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

func checkLine(sudoku [9][9]int, n, row, column int) bool {
	for i := 0; i < 9; i++ {
		if i == row {
			continue
		}

		if sudoku[i][column] == n {
			return false
		}
	}

	return true
}

func checkColumn(sudoku [9][9]int, n, row, column int) bool {
	for i := 0; i < 9; i++ {
		if i == column {
			continue
		}

		if sudoku[row][i] == n {
			return false
		}
	}

	return true
}

func getBounds(axis int) (int, int) {
	if axis < 3 {
		return 0, 2
	} else if axis < 6 {
		return 3, 5
	} else {
		return 6, 8
	}
}
