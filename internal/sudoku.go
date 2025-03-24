package internal

import (
	"fmt"
	"time"
)

func Sudoku(sudokuStr string) [9][9]int8 {
	elapsedTime := time.Now()

	sudoku := parse(sudokuStr)

	if VerifySudoku(sudoku) {
		sudoku = resolve(sudoku)
		printSudoku(sudoku)
	} else {
		panic("Impossible sudoku")
	}

	fmt.Println("Time:", time.Now().Sub(elapsedTime).Milliseconds(), "ms")
	return sudoku
}
