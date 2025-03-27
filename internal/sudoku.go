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
