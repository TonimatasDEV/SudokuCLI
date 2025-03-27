package main

import (
	"SudokuCLI/actions"
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
		case "select":
			if len(command) < 2 {
				fmt.Println("Please enter de sudoku with as argument.")
			}

			actions.SelectSudoku(args[0])
			fmt.Println("What do you want to do? If you don't know, use \"help\" command.")
			break
		case "play":
			actions.PlaySudoku(scanner)
			fmt.Println("What do you want to do now?")
			break
		case "resolve":
			actions.ResolveSudoku()
			fmt.Println("Easy one! Now, what do you want to do?")
			break
		case "generate":
			actions.GenerateSudoku()
			fmt.Println("I think its good one. What do you want to do with it?")
			break
		case "export":
			actions.ExportSudoku()
			fmt.Println("Here is! What do you want to do?")
			break
		case "help":
			fmt.Printf("Commands:\n" +
				" - select <sudoku>  Select a sudoku. Format: firstRow,secondRow,thirdRow,...\n" +
				" - play             Play the current selected sudoku.\n" +
				" - resolve          Resolve the current selected sudoku.\n" +
				" - generate         Generates a new random Sudoku and select it.\n" +
				" - exit             Exit the program.\n")
			break
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Unknown command. Use \"help\" command to know what commands exist.")
		}
	}
}
