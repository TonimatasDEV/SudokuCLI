package internal

import (
	"math/rand"
)

func CheckNumber(sudoku [9][9]int, num, row, column int) bool {
	for i := 0; i < 9; i++ {
		if sudoku[row][i] == num || sudoku[i][column] == num {
			return false
		}
	}

	startRow := row / 3 * 3
	startColumn := column / 3 * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if sudoku[startRow+i][startColumn+j] == num {
				return false
			}
		}
	}

	return true
}

func FillSudoku(sudoku *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if sudoku[row][column] == 0 {
				nums := rand.Perm(9)

				for _, num := range nums {
					if CheckNumber(*sudoku, num+1, row, column) {
						sudoku[row][column] = num + 1
						if FillSudoku(sudoku) {
							return true
						}
						sudoku[row][column] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func RemoveNumbers(sudoku *[9][9]int, count int) {
	for i := 0; i < count; {
		row, col := rand.Intn(9), rand.Intn(9)
		if sudoku[row][col] != 0 {
			backup := sudoku[row][col]
			sudoku[row][col] = 0

			copySudoku := *sudoku
			solutions := 0
			SolveSudoku(&copySudoku, &solutions)
			if solutions != 1 {
				sudoku[row][col] = backup
			} else {
				i++
			}
		}
	}
}

func SolveSudoku(sudoku *[9][9]int, solutions *int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if sudoku[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if CheckNumber(*sudoku, num, row, col) {
						sudoku[row][col] = num

						SolveSudoku(sudoku, solutions)

						if *solutions > 1 {
							return
						}

						sudoku[row][col] = 0
					}
				}

				return
			}
		}
	}

	*solutions++
}
