package internal

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(sudokuStr string) [9][9]int8 {
	var sudoku [9][9]int8

	for i, line := range strings.Split(sudokuStr, ",") {
		for k, char := range line {
			integer, err := strconv.Atoi(string(char))

			if err != nil {
				panic("Invalid format")
			}

			sudoku[i][k] = int8(integer)
		}
	}

	return sudoku
}

func printSudoku(sudoku [9][9]int8) {
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
