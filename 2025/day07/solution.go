package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	// P1()
	P2()
}

type coord struct {
	Row, Col int
}

func (c coord) Eq(other coord) bool {
	return c.Row == other.Row && c.Col == other.Col
}

func (c coord) String() string {
	return fmt.Sprintf("%d:%d", c.Row, c.Col)
}

func P2() {
	data, err := os.ReadFile("assets/input.txt")
	if err != nil {
		panic(err)
	}

	rows := bytes.Split(bytes.TrimSpace(data), []byte{'\n'})
	row := len(rows) - 1
	branches := 0
	for col := range rows[row] {
		branches += possiblePaths(coord{row, col}, rows)
	}
	fmt.Println(branches)
}

var cache = make(map[string]int)

func possiblePaths(c coord, rows [][]byte) int {
	if num, ok := cache[c.String()]; ok {
		return num
	}
	var num int
	rowOffset := 1
	for {
		center := coord{c.Row - rowOffset, c.Col}
		cChar := rows[center.Row][center.Col]
		if center.Row == 0 {
			if cChar == 'S' {
				num = 1
			}
			break
		}
		if cChar == '^' {
			break
		}
		if c.Col >= 1 {
			left := coord{c.Row - rowOffset, c.Col - 1}
			if rows[left.Row][left.Col] == '^' {
				num += possiblePaths(left, rows)
			}
		}
		if c.Col < len(rows[0])-1 {
			right := coord{c.Row - rowOffset, c.Col + 1}
			if rows[right.Row][right.Col] == '^' {
				num += possiblePaths(right, rows)
			}
		}
		rowOffset++
	}
	cache[c.String()] = num
	return num
}

func P1() {
	data, err := os.ReadFile("assets/input.txt")
	if err != nil {
		panic(err)
	}

	var stack []coord
	rows := bytes.Split(bytes.TrimSpace(data), []byte{'\n'})
	for col, char := range rows[0] {
		if char == 'S' {
			stack = append(stack, coord{Row: 0, Col: col})
			break
		}
	}
	var split int
	for len(stack) > 0 {
		var tip coord
		tip, stack = stack[0], stack[1:]
		if tip.Row == len(rows)-1 {
			// in the bottom row, nothing to do
			continue
		}
		switch rows[tip.Row+1][tip.Col] {
		case '.':
			newCoord := coord{Row: tip.Row + 1, Col: tip.Col}
			if !slices.ContainsFunc(stack, newCoord.Eq) {
				stack = append(stack, newCoord)
			}
		case '^':
			split++

			left := coord{tip.Row + 1, tip.Col - 1}
			if !slices.ContainsFunc(stack, left.Eq) {
				stack = append(stack, left)
			}

			right := coord{tip.Row + 1, tip.Col + 1}
			if !slices.ContainsFunc(stack, right.Eq) {
				stack = append(stack, right)
			}
		}
	}
	fmt.Println(split)
}

func printBoard(rows [][]byte) {
	for _, row := range rows {
		fmt.Println(string(row))
	}
}
