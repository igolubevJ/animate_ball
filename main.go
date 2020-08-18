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

	buf := make([]rune, 0, width * height)

	buf = buf[:0]

	// draw the board
	for y := range board[0] {
		for x := range board {
			cell = cellEmpty
			if board[x][y] {
				cell = cellBall
			}
			buf = append(buf, cell, ' ')
		}

		buf = append(buf, '\n')
	}
	fmt.Print(string(buf))
}
