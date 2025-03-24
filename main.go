package main

import (
	"SudokuCLI/internal"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: sudokucli <sudoku>")
		fmt.Println("Sudoku format:")
		fmt.Println("firstRow,secondRow,thirdRow...")
		fmt.Println("Example: sudokucli \"100504030,206000400...\"")
		os.Exit(1)
	}

	internal.Sudoku(args[1])
}
