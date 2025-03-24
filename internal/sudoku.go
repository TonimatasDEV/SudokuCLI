package internal

import (
	"fmt"
	"time"
)

var sudoku [9][9]int

func Sudoku(sudoku string) {
	elapsedTime := time.Now()

	parse(sudoku)

	finishedSudoku := resolve()
	printSudoku(finishedSudoku)

	fmt.Println("Time:", time.Now().Sub(elapsedTime).Milliseconds(), "ms")
}
