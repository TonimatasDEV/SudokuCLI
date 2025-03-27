package actions

import (
	"SudokuCLI/internal"
	"fmt"
	"strconv"
	"time"
)

var selectedSudoku [9][9]int

func SelectSudoku(sudokuStr string) {
	startTime := time.Now()

	if !internal.Parse(&selectedSudoku, sudokuStr) {
		selectedSudoku = [9][9]int{}
		internal.PrintlnWithElapsedTime("Invalid sudoku format.", startTime)
	}

	solutions := 0
	internal.SolveSudoku(&selectedSudoku, &solutions)

	if solutions > 1 || solutions == 0 {
		selectedSudoku = [9][9]int{}
		internal.PrintlnWithElapsedTime("This sudoku is impossible.", startTime)
	} else {
		internal.PrintlnWithElapsedTime("Sudoku parsed and selected.", startTime)
		internal.PrintSudoku(selectedSudoku)
	}
}

func ResolveSudoku() {
	startTime := time.Now()

	internal.FillSudoku(&selectedSudoku)

	internal.PrintlnWithElapsedTime("Sudoku resolved.", startTime)
	internal.PrintSudoku(selectedSudoku)
}

func GenerateSudoku() {
	startTime := time.Now()

	internal.FillSudoku(&selectedSudoku)
	internal.RemoveNumbers(&selectedSudoku, 30)

	internal.PrintlnWithElapsedTime("Generated successfully.", startTime)
	internal.PrintSudoku(selectedSudoku)
}

func ExportSudoku() {
	startTime := time.Now()

	var str string
	for column := 0; column < 9; column++ {
		for row := 0; row < 9; row++ {
			str += strconv.Itoa(int(selectedSudoku[row][column]))
		}

		str += ","
	}

	fmt.Println("Sudoku:", str)
	internal.PrintlnWithElapsedTime("Sudoku exported.", startTime)
}
