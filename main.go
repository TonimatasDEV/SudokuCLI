package main

import (
	"SudokuCLI/src"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to SudokuCLI v1.0\nYou don't know how to use it? Use \"help\" command.")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		input := scanner.Text()
		rawSplit := strings.Split(input, " ")
		command := rawSplit[0]
		args := rawSplit[1:]

		switch command {
		case "play":
			src.PlaySudoku(scanner)
			fmt.Println("What do you want to do now?")
			break
		case "resolve":
			src.ResolveSudoku()
			fmt.Println("Easy one! Now, what do you want to do?")
			break
		case "generate":
			src.GenerateSudoku()
			fmt.Println("I think its good one. What do you want to do with it?")
			break
		case "export":
			src.ExportSudoku()
			fmt.Println("Here is! What do you want to do?")
			break
		case "import":
			if len(command) < 2 {
				fmt.Println("Please enter the sudoku.")
			}

			src.ImportSudoku(args[0])
			fmt.Println("What do you want to do? If you don't know, use \"help\" command.")
			break
		case "help":
			fmt.Println("Commands:\n" +
				" - help             This message.\n" +
				" - generate         Generates a new random Sudoku and select it.\n" +
				" - play             Play the current selected sudoku.\n" +
				" - resolve          Resolve the current selected sudoku.\n" +
				" - export           Export the sudoku with the program format.\n" +
				" - import <sudoku>  Import a sudoku. Format: firstRow,secondRow,thirdRow,...\n" +
				" - exit             Exit the program.")
			break
		case "exit":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Unknown command. Use \"help\" command to know what commands exist.")
		}
	}
}
