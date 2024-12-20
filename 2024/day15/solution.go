package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	robot, board, directions := puzzle("assets/input.txt")
	b2, r2 := board.Scale()
	p1(robot, board, directions)
	p2(r2, b2, directions)
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

func p2(robot *Robot, board Board, directions []Direction) {
	for _, direction := range directions {
		robot.Move(direction, board)
	}
	var total int
	for y, row := range board {
		for x, col := range row {
			if _, ok := col.(*BoxLeft); ok {
				total += (100 * y) + x
			}
		}
	}
	fmt.Println(total)
}

type Mover interface {
	CanMove(d Direction, board Board) bool
	Move(d Direction, board Board)
	String() string
}

type Board [][]Mover

func (b Board) String() string {
	var sb strings.Builder
	for _, row := range b {
		for _, col := range row {
			var s string
			if col == nil {
				s = "."
			} else {
				s = col.String()
			}
			fmt.Fprint(&sb, s)
		}
		fmt.Fprint(&sb, "\n")
	}
	return sb.String()
}

func (b Board) Scale() (Board, *Robot) {
	out := make(Board, len(b))
	for i := range out {
		out[i] = make([]Mover, len(b[0])*2)
	}
	var r Robot
	for y, row := range b {
		for x, col := range row {
			nx := x * 2
			switch col.(type) {
			case *Box:
				l := BoxLeft{X: nx, Y: y}
				r := BoxRight{X: nx + 1, Y: y}

				l.Right = &r
				r.Left = &l

				out[y][nx] = &l
				out[y][nx+1] = &r
			case *Wall:
				out[y][nx] = col
				out[y][nx+1] = col
			case *Robot:
				r.Y = y
				r.X = nx
				out[y][nx] = &r
			}
		}
	}
	return out, &r
}

type Wall struct{}

func (w *Wall) CanMove(Direction, Board) bool { return false }
func (w *Wall) Move(Direction, Board)         { panic("cannot move wall") }
func (w *Wall) String() string                { return "#" }

type Box struct {
	X, Y int
}

func (b *Box) String() string { return "0" }
func (b *Box) CanMove(d Direction, board Board) bool {
	deltas := dirDeltas[d]
	nx, ny := b.X+deltas[0], b.Y+deltas[1]
	n := board[ny][nx]
	if n == nil {
		return true
	}
	return n.CanMove(d, board)
}

func (b *Box) Move(d Direction, board Board) {
	deltas := dirDeltas[d]
	nx, ny := b.X+deltas[0], b.Y+deltas[1]
	n := board[ny][nx]
	if n == nil {
		board[b.Y][b.X] = nil
		board[ny][nx] = b
		b.X, b.Y = nx, ny
	} else if n.CanMove(d, board) {
		n.Move(d, board)
		board[b.Y][b.X] = nil
		board[ny][nx] = b
		b.X, b.Y = nx, ny
	}
}

type Robot struct {
	X, Y int
}

func (r *Robot) String() string { return "@" }
func (r *Robot) CanMove(d Direction, board Board) bool {
	deltas := dirDeltas[d]
	nx, ny := r.X+deltas[0], r.Y+deltas[1]
	n := board[ny][nx]
	if n == nil {
		return true
	}
	return n.CanMove(d, board)
}

func (r *Robot) Move(d Direction, board Board) {
	deltas := dirDeltas[d]
	nx, ny := r.X+deltas[0], r.Y+deltas[1]
	n := board[ny][nx]
	if n == nil {
		board[r.Y][r.X] = nil
		board[ny][nx] = r
		r.X, r.Y = nx, ny
	} else if n.CanMove(d, board) {
		n.Move(d, board)
		board[r.Y][r.X] = nil
		board[ny][nx] = r
		r.X, r.Y = nx, ny
	}
}

type BoxLeft struct {
	X, Y  int
	Right *BoxRight
}

func (b *BoxLeft) String() string { return "[" }
func (b *BoxLeft) canMove(d Direction, board Board) bool {
	if d == right {
		return true
	}
	deltas := dirDeltas[d]
	nx, ny := b.X+deltas[0], b.Y+deltas[1]
	n := board[ny][nx]
	if n == nil {
		return true
	}
	return n.CanMove(d, board)
}

func (b *BoxLeft) CanMove(d Direction, board Board) bool {
	return b.canMove(d, board) && b.Right.canMove(d, board)
}

func (b *BoxLeft) move(d Direction, board Board) {
	deltas := dirDeltas[d]
	nx, ny := b.X+deltas[0], b.Y+deltas[1]
	n := board[ny][nx]
	if n == nil {
		board[b.Y][b.X] = nil
		board[ny][nx] = b
		b.X, b.Y = nx, ny
	} else if n != b.Right && n.CanMove(d, board) {
		n.Move(d, board)
		board[b.Y][b.X] = nil
		board[ny][nx] = b
		b.X, b.Y = nx, ny
	}
}

func (b *BoxLeft) Move(d Direction, board Board) {
	if d == right {
		// if moving right, make sure we move the right box first
		b.Right.move(d, board)
		b.move(d, board)
	} else {
		b.move(d, board)
		b.Right.move(d, board)
	}
}

type BoxRight struct {
	X, Y int
	Left *BoxLeft
}

func (b *BoxRight) String() string { return "]" }

func (b *BoxRight) CanMove(d Direction, board Board) bool {
	return b.canMove(d, board) && b.Left.canMove(d, board)
}

func (b *BoxRight) canMove(d Direction, board Board) bool {
	if d == left {
		return true
	}
	deltas := dirDeltas[d]
	nx, ny := b.X+deltas[0], b.Y+deltas[1]
	n := board[ny][nx]
	if n == nil {
		return true
	}
	return n.CanMove(d, board)
}

func (b *BoxRight) move(d Direction, board Board) {
	deltas := dirDeltas[d]
	nx, ny := b.X+deltas[0], b.Y+deltas[1]
	n := board[ny][nx]
	if n == nil {
		board[b.Y][b.X] = nil
		board[ny][nx] = b
		b.X, b.Y = nx, ny
	} else if n != b.Left && n.CanMove(d, board) {
		n.Move(d, board)
		board[b.Y][b.X] = nil
		board[ny][nx] = b
		b.X, b.Y = nx, ny
	}
}

func (b *BoxRight) Move(d Direction, board Board) {
	if d == left {
		// if moving left, make sure we move the left box first
		b.Left.move(d, board)
		b.move(d, board)
	} else {
		b.move(d, board)
		b.Left.move(d, board)
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
		b[y] = make([]Mover, len(row))
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
