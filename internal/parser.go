package internal

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(sudoku *[9][9]int, sudokuStr string) bool {
	for i, line := range strings.Split(sudokuStr, ",") {
		for k, char := range line {
			integer, err := strconv.Atoi(string(char))

			if err != nil {
				return false
			}

			sudoku[i][k] = integer
		}
	}

	return true
}

func PrintSudoku(sudoku [9][9]int) {
	fmt.Println("-------------------------")
	for row := range sudoku {
		for column := range sudoku[row] {
			switch column {
			case 0:
				fmt.Print("| ", sudoku[row][column], " ")
				break
			case 2, 5:
				fmt.Print(sudoku[row][column], " | ")
				break
			case 8:
				fmt.Print(sudoku[row][column], " |")
				break
			default:
				fmt.Print(sudoku[row][column], " ")
				break
			}
		}

		if row == 2 || row == 5 {
			fmt.Println("\n|-----------------------|")
		} else {
			fmt.Println()
		}
	}

	fmt.Print("-------------------------\n")
}
