package internal

import (
	"fmt"
	"time"
)

var selectedSudoku [9][9]int8

func SelectSudoku(sudokuStr string) {
	startTime := time.Now()

	sudoku := parse(sudokuStr)

	if !VerifySudoku(sudoku) {
		printlnWithElapsedTime("This sudoku is impossible.", startTime)
	} else {
		selectedSudoku = sudoku
		printlnWithElapsedTime("Sudoku parsed and selected.", startTime)
		printSudoku(selectedSudoku)
		fmt.Println("What do you want to do? If you don't know, use \"help\" command.")
	}
}

func ResolveSudoku() {
	startTime := time.Now()
	selectedSudoku = resolve(selectedSudoku)
	printlnWithElapsedTime("Sudoku resolved.", startTime)
	printSudoku(selectedSudoku)
}

func GenerateSudoku() {
	startTime := time.Now()
	var sudoku [9][9]int8

	for i := 0; i < 30; i++ {
		number := randomNum(9)
		row := randomNum(8)
		column := randomNum(8)

		if sudoku[row][column] == 0 {
			newSudoku, ok := addNumber(sudoku, number, row, column)

			if ok && VerifySudoku(newSudoku) {
				sudoku = newSudoku
				continue
			}
		}

		i--
	}

	selectedSudoku = sudoku
	printlnWithElapsedTime("", startTime)
	printSudoku(selectedSudoku)
}
