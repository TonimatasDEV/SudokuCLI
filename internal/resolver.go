package internal

func resolve() [9][9]int {
	finishedSudoku, ok := tryValidNumbers(sudoku, 0, 0)

	if ok {
		return finishedSudoku
	} else {
		panic("Sudoku is invalid")
	}
}

func tryValidNumbers(sudoku [9][9]int, row, column int) ([9][9]int, bool) {
	validNumbers := whatNumbersArePossible(sudoku, row, column)
	nextRow := row
	nextColumn := column + 1

	if nextColumn > 8 {
		nextColumn = 0
		nextRow++
	}

	if row > 8 {
		return sudoku, true
	}

	if sudoku[row][column] != 0 {
		sudoku, ok := tryValidNumbers(sudoku, nextRow, nextColumn)
		if ok {
			return sudoku, true
		}
	} else {
		for _, value := range validNumbers {
			sudoku, ok := addNumber(sudoku, value, row, column)
			if ok {
				sudoku, ok1 := tryValidNumbers(sudoku, nextRow, nextColumn)
				if ok1 {
					return sudoku, true
				}
			}
		}
	}

	return sudoku, false
}

func whatNumbersArePossible(sudoku [9][9]int, row, column int) []int {
	var numbers []int

	for i := 1; i <= 9; i++ {
		if checkNumber(sudoku, i, row, column) {
			numbers = append(numbers, i)
		}
	}

	return numbers
}

func addNumber(sudoku [9][9]int, n, row, column int) ([9][9]int, bool) {
	if checkNumber(sudoku, n, row, column) {
		sudoku[row][column] = n
		return sudoku, true
	}

	return sudoku, false
}
