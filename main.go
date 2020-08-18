package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		width = 50
		height = 10

		cellEmpty = ' '
		cellBall = '🏀'

		maxFrames = 1200

		speed = time.Second / 20
	)

	var (
		px, py int
		vx, vy = 1, 1

		cell rune
	)

	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	for i := 0; i < maxFrames; i++ {
		px += vx
		py += vy

		if px <= 0 || px >= height - 1 {
			vx *= -1
		}

		if py <= 0 || py >= width - 1 {
			vy *= -1
		}

		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		board[py][px] = true  

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
		
		moveTopLeft()

		fmt.Print(string(buf))

		time.Sleep(speed)
	}
}

func moveTopLeft() {
	fmt.Print("\033[H")
}