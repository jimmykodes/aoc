package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fn := "assets/input.txt"
	p1(fn)
	p2(fn)
}

func p1(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	rows := bytes.Split(data, []byte{'\n'})
	numRows := len(rows)
	numCols := len(rows[0])
	count := 0
	for y := range numRows {
		for x := range numCols {
			if rows[y][x] == 'X' {
				left := x >= 3
				right := x <= numCols-4
				up := y >= 3
				down := y <= numRows-4
				if left && check(rows, x, y, -1, 0) {
					count++
				}
				if right && check(rows, x, y, 1, 0) {
					count++
				}
				if up && check(rows, x, y, 0, -1) {
					count++
				}
				if down && check(rows, x, y, 0, 1) {
					count++
				}
				if left && up && check(rows, x, y, -1, -1) {
					count++
				}
				if left && down && check(rows, x, y, -1, 1) {
					count++
				}
				if right && up && check(rows, x, y, 1, -1) {
					count++
				}
				if right && down && check(rows, x, y, 1, 1) {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func check(board [][]byte, x, y, dx, dy int) bool {
	xmas := []byte{'X', 'M', 'A', 'S'}
	for i := range 4 {
		if board[y+(dy*i)][x+(dx*i)] != xmas[i] {
			return false
		}
	}
	return true
}

func checkX(board [][]byte, x, y int) bool {
	upperLeft := board[y-1][x-1]
	upperRight := board[y-1][x+1]
	lowerLeft := board[y+1][x-1]
	lowerRight := board[y+1][x+1]

	leftMas := (upperLeft == 'M' && lowerRight == 'S') || (upperLeft == 'S' && lowerRight == 'M')
	rightMas := (upperRight == 'M' && lowerLeft == 'S') || (upperRight == 'S' && lowerLeft == 'M')

	return leftMas && rightMas
}

func p2(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	rows := bytes.Split(data, []byte{'\n'})
	numRows := len(rows)
	numCols := len(rows[0])
	count := 0
	for y := range numRows {
		for x := range numCols {
			if rows[y][x] == 'A' {
				left := x >= 1
				right := x <= numCols-2
				up := y >= 1
				down := y <= numRows-2
				if left && right && up && down && checkX(rows, x, y) {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
