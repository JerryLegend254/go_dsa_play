package main

import "fmt"

const N = 9

func printBoard(board *[N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
}

// Function to check if placing a number is valid
func isValid(board *[N][N]int, row, col, num int) bool {
	// Check row
	for x := 0; x < N; x++ {
		if board[row][x] == num {
			return false
		}
	}

	// Check column
	for x := 0; x < N; x++ {
		if board[x][col] == num {
			return false
		}
	}

	// Check subgrid
	startRow, startCol := row-row%3, col-col%3 //if subgrid is 3x3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

// Backtracking function to solve the board
func solveSudoku(board *[N][N]int) bool {
	row, col := -1, -1
	foundEmpty := false

	// Find an empty cell
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 0 {
				row, col = i, j
				foundEmpty = true
				break
			}
		}
		if foundEmpty {
			break
		}
	}

	// If all non-empty cells then already solved
	if !foundEmpty {
		return true
	}

	// Try placing numbers 1 to 9
	for num := 1; num <= 9; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if solveSudoku(board) {
				return true
			}
			// Backtrack
			board[row][col] = 0
		}
	}

	return false
}

func main() {
	board := [N][N]int{
		{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}

	//testing with an example
	if solveSudoku(&board) {
		printBoard(&board)
	} else {
		fmt.Println("No solution exists")
	}
}
