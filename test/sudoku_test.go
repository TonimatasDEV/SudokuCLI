package test

import (
	"SudokuCLI/internal"
	"fmt"
	"strings"
	"testing"
	"time"
)

var sudokus = []string{
	"906000000,020005070,000002400,000039000,307000010,000000006,700800205,405600030,100300000",
	"000803000,500000003,040000020,006000500,080105070,007000600,030000080,400000007,000209000-1",
	"800000000,003600000,070090200,050007000,000045700,000100030,001000068,008500010,090000400",
	"020060000,000007009,600400800,100000000,000030000,000000007,003008002,700200000,000010080-1",
	"000000907,000420180,000705026,100904000,050000040,000507009,920108000,034059000,507000000",
}

func TestSudoku(t *testing.T) {
	for i, str := range sudokus {
		var sudoku [9][9]int
		var solutions int

		strSplit := strings.Split(str, "-")

		if !internal.Parse(&sudoku, strSplit[0]) {
			t.Error("Incorrect parse")
		}

		internal.SolveSudoku(&sudoku, &solutions)

		if solutions > 1 {
			if len(strSplit) < 2 {
				t.Error("Incorrect solutions")
			} else {
				if strSplit[1] == "1" {
					fmt.Println("Sudoku", i+1, "has more than one solution.")
					continue
				} else {
					t.Error("Incorrect solutions")
				}
			}
		}

		if !internal.FillSudoku(&sudoku) {
			t.Error("Incorrect fill")
		}

		internal.PrintSudoku(sudoku)
	}
}

func TestInvalidSudokuStr(t *testing.T) {
	var sudoku [9][9]int

	if internal.Parse(&sudoku, "1-2") {
		t.Error("Incorrect parse")
	}
}

func TestElapsedTime(t *testing.T) {
	startTime := time.Now()

	time.Sleep(25 * time.Millisecond)

	internal.PrintlnWithElapsedTime("Time:", startTime)
}

var remove_sudokus = []string{
	"639875412,475162389,218394765,127538694,896427153,354619278,961253847,743981526,582746931",
	"368792415,142365897,795841236,853974162,916258374,427136589,279483651,681527943,534619728",
	"843165792,926784315,175932648,497216583,231458967,568379124,789521436,654893271,312647859",
	"219564378,758932146,436178295,967421853,342785619,581396724,625819437,874253961,193647582",
	"625193784,738642159,914857623,592364871,467981532,183725946,241576398,876239415,359418267",
}

func TestRemoveNumbers(t *testing.T) {
	for _, str := range remove_sudokus {
		var sudoku [9][9]int

		if !internal.Parse(&sudoku, str) {
			t.Error("Incorrect parse")
		}

		internal.RemoveNumbers(&sudoku, 50)

		var solutions int

		internal.SolveSudoku(&sudoku, &solutions)

		if solutions > 1 {
			t.Error("Incorrect remove numbers")
		}

		internal.PrintSudoku(sudoku)
	}

}
