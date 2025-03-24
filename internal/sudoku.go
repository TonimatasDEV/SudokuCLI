package internal

import (
	"fmt"
	"time"
)

var sudoku [9][9]int8

func Sudoku(sudokuStr string) {
	elapsedTime := time.Now()

	parse(sudokuStr)

	finishedSudoku := resolve()
	printSudoku(finishedSudoku)

	fmt.Println("Time:", time.Now().Sub(elapsedTime).Milliseconds(), "ms")
}
