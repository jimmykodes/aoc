package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	robot, board, directions := puzzle("assets/input.txt")
	p1(robot, board, directions)
}

func p1(robot *Robot, board Board, directions []Direction) {
	for _, direction := range directions {
		robot.Move(direction, board)
	}
	var total int
	for y, row := range board {
		for x, col := range row {
			if _, ok := col.(*Box); ok {
				total += (100 * y) + x
			}
		}
	}
	fmt.Println(total)
}

type Board [][]any

func (b Board) String() string {
	var sb strings.Builder
	for _, row := range b {
		for _, col := range row {
			var s string
			switch col.(type) {
			case *Box:
				s = "O"
			case *Wall:
				s = "#"
			case *Robot:
				s = "@"
			default:
				s = "."
			}
			fmt.Fprint(&sb, s)
		}
		fmt.Fprint(&sb, "\n")
	}
	return sb.String()
}

type (
	Wall struct{}
	Box  struct {
		X, Y int
	}
	Robot struct {
		X, Y int
	}
)

func (b *Box) Move(d Direction, board Board) bool {
	deltas := dirDeltas[d]
	nx, ny := b.X+deltas[0], b.Y+deltas[1]
	switch n := board[ny][nx].(type) {
	case *Box:
		if !n.Move(d, board) {
			return false
		}
		board[b.Y][b.X] = nil
		board[ny][nx] = b
		b.X, b.Y = nx, ny
		return true
	case *Wall:
		return false
	case nil:
		board[b.Y][b.X] = nil
		board[ny][nx] = b
		b.X, b.Y = nx, ny
		return true
	default:
		panic(fmt.Sprintf("box move: invalid type: %T", n))
	}
}

func (r *Robot) Move(d Direction, board Board) bool {
	deltas := dirDeltas[d]
	nx, ny := r.X+deltas[0], r.Y+deltas[1]
	switch n := board[ny][nx].(type) {
	case *Box:
		if !n.Move(d, board) {
			return false
		}
		board[r.Y][r.X] = nil
		r.X, r.Y = nx, ny
		board[ny][nx] = r
		return true
	case *Wall:
		return false
	case nil:
		board[r.Y][r.X] = nil
		board[ny][nx] = r
		r.X, r.Y = nx, ny
		return true
	default:
		panic(fmt.Sprintf("robot move: invalid type: %T", n))
	}
}

type Direction int

func (d Direction) String() string {
	return [...]string{
		up:    "^",
		down:  "v",
		right: ">",
		left:  "<",
	}[d]
}

const (
	up Direction = iota
	right
	down
	left
)

var dirDeltas = [...][2]int{
	up:    {0, -1},
	right: {1, 0},
	down:  {0, 1},
	left:  {-1, 0},
}

func puzzle(filename string) (*Robot, Board, []Direction) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	boardData, movementData, _ := bytes.Cut(data, []byte{'\n', '\n'})
	rowData := bytes.Split(boardData, []byte{'\n'})
	movementData = bytes.Replace(movementData, []byte{'\n'}, []byte{}, -1)

	var r *Robot
	b := make(Board, len(rowData))
	for y, row := range rowData {
		b[y] = make([]any, len(row))
		for x, col := range row {
			switch col {
			case '#':
				b[y][x] = &Wall{}
			case 'O':
				b[y][x] = &Box{X: x, Y: y}
			case '@':
				r = &Robot{X: x, Y: y}
				b[y][x] = r
			}
		}
	}
	d := make([]Direction, len(movementData))
	for i, char := range movementData {
		switch char {
		case '^':
			d[i] = up
		case '>':
			d[i] = right
		case 'v':
			d[i] = down
		case '<':
			d[i] = left
		}
	}
	return r, b, d
}
