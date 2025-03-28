package src

import (
	"SudokuCLI/internal"
	"bufio"
	"fmt"
	"strconv"
	"strings"
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

func PlaySudoku(scanner *bufio.Scanner) {
	resolvedSudoku := selectedSudoku

	internal.FillSudoku(&resolvedSudoku)

	fmt.Println("Playing Sudoku...\n" +
		"When I ask you, I want you to respond with the number you want to place and where, using the following format.\n" +
		"If you want to exit, only type exit.\n" +
		"Format: number,row,column - e.g. 5,2,1")

	for {
		if resolvedSudoku == selectedSudoku {
			fmt.Println("Incredible! You finished the sudoku!")
			break
		}

		internal.PrintSudoku(selectedSudoku)
		fmt.Println("What number do you want to place in the Sudoku?")

		scanner.Scan()
		input := scanner.Text()
		rawSplit := strings.Split(input, ",")

		if rawSplit[0] == "exit" {
			fmt.Println("Why? You didn't finish the sudoku.")
			break
		}

		if len(rawSplit) != 3 {
			fmt.Println("Please use the requested format.")
			continue
		}

		number, err1 := strconv.Atoi(rawSplit[0])
		row, err2 := strconv.Atoi(rawSplit[1])
		column, err3 := strconv.Atoi(rawSplit[2])

		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Please use the requested format.")
			continue
		}

		if resolvedSudoku[row-1][column-1] == number {
			selectedSudoku[row-1][column-1] = number
			fmt.Println("Nice!")
		} else {
			fmt.Println(fmt.Sprintf("Oh no! %d is incorrect number for row %d and column %d.", number, row, column))
		}
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

	selectedSudoku = [9][9]int{}
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

		if column < 8 {
			str += ","
		}
	}

	fmt.Println("Sudoku:", str)
	internal.PrintlnWithElapsedTime("Sudoku exported.", startTime)
}
