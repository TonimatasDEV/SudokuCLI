package test

import (
	"SudokuCLI/internal"
	"testing"
)

func TestSudoku(t *testing.T) {
	internal.Sudoku("906000000,020005070,000002400,000039000,307000010,000000006,700800205,405600030,100300000")
}
