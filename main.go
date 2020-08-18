package main

import "fmt"

func main() {
	const (
		width = 50
		height = 10

		cellEmpty = ' '
		cellBall = '🏀'
	)

	var cell rune

	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	board[0][0] = true

	// draw the board
	for y := range board[0] {
		for x := range board {
			cell = cellEmpty
			if board[x][y] {
				cell = cellBall
			}
			fmt.Print(string(cell), " ")
		}

		fmt.Println()
	}
}
