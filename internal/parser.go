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
			str := strconv.Itoa(sudoku[row][column])

			if str == "0" {
				str = "·"
			}

			switch column {
			case 0:
				fmt.Print("| ", str, " ")
				break
			case 2, 5:
				fmt.Print(str, " | ")
				break
			case 8:
				fmt.Print(str, " |")
				break
			default:
				fmt.Print(str, " ")
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
