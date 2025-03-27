package internal

func resolve(sudoku [9][9]int8) [9][9]int8 {
	finishedSudoku, ok := tryValidNumbers(sudoku, 0, 0)

	if ok {
		return finishedSudoku
	} else {
		panic("Impossible sudoku")
	}
}

func tryValidNumbers(sudoku [9][9]int8, row, column int8) ([9][9]int8, bool) {
	if row > 8 {
		return sudoku, true
	}

	validNumbers := whatNumbersArePossible(sudoku, row, column)
	nextRow := row
	nextColumn := column + 1

	if nextColumn > 8 {
		nextColumn = 0
		nextRow++
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
				sudoku, ok := tryValidNumbers(sudoku, nextRow, nextColumn)
				if ok {
					return sudoku, true
				}
			}
		}
	}

	return sudoku, false
}

func whatNumbersArePossible(sudoku [9][9]int8, row, column int8) []int8 {
	var numbers []int8

	var i int8
	for i = 1; i <= 9; i++ {
		if checkNumber(sudoku, i, row, column) {
			numbers = append(numbers, i)
		}
	}

	return numbers
}

func addNumber(sudoku [9][9]int8, n, row, column int8) ([9][9]int8, bool) {
	if checkNumber(sudoku, n, row, column) {
		sudoku[row][column] = n
		return sudoku, true
	}

	return sudoku, false
}
