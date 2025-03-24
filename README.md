# SudokuCLI
SudokuCLI is a program to resolve Sudokus.

# How it works?
1. Clone the repository
2. Execute `go build` to compile the program.
3. Run the program with the sudoku. Example `SudokuCLI "firstRow,secoundRow,thirdRow..."`

# Execute example
Command: `.\SudokuCLI.exe "906000000,020005070,000002400,000039000,307000010,000000006,700800205,405600030,100300000"`

- Output:
```
-------------------------
| 9 7 6 | 4 1 3 | 5 8 2 |
| 8 2 4 | 9 6 5 | 1 7 3 |
| 5 3 1 | 7 8 2 | 4 6 9 |
|-----------------------|
| 6 1 8 | 5 3 9 | 7 2 4 |
| 3 5 7 | 2 4 6 | 9 1 8 |
| 2 4 9 | 1 7 8 | 3 5 6 |
|-----------------------|
| 7 6 3 | 8 9 1 | 2 4 5 |
| 4 9 5 | 6 2 7 | 8 3 1 |
| 1 8 2 | 3 5 4 | 6 9 7 |
-------------------------
Time: X ms
```