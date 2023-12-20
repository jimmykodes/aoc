package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	board := load("assets/input.txt")
	fmt.Println(p1(board))
	fmt.Println(p2(board))
}

type Direction int

const (
	North Direction = iota + 1
	East
	South
	West
)

type Light struct {
	Y, X int
	Dir  Direction
}

func (l *Light) Move() {
	switch l.Dir {
	case North:
		l.Y--
	case East:
		l.X++
	case South:
		l.Y++
	case West:
		l.X--
	}
}

type Tile interface {
	Interact(*Light) (*Light, *Light)
	IsEnergized() bool
	Clear()
}

type BaseTile struct {
	energized bool
	dirs      [5]bool
}

func (b BaseTile) IsEnergized() bool {
	return b.energized
}

func (b *BaseTile) Clear() {
	b.energized = false
	b.dirs = [5]bool{}
}

func (t *BaseTile) Seen(l *Light) bool {
	t.energized = true
	if t.dirs[l.Dir] {
		return true
	}
	t.dirs[l.Dir] = true
	return false
}

type EmptyTile struct{ BaseTile }

func (t *EmptyTile) Interact(l *Light) (*Light, *Light) {
	t.energized = true
	l.Move()
	return l, nil
}

type HSplitter struct{ BaseTile }

func (s *HSplitter) Interact(l *Light) (*Light, *Light) {
	if s.Seen(l) {
		return nil, nil
	}
	if l.Dir == North || l.Dir == South {
		// a horizontal splitter takes North or South light and splits
		// it to east & west light
		l2 := *l
		l.Dir = West
		l2.Dir = East
		l.Move()
		l2.Move()
		return l, &l2
	}
	l.Move()
	return l, nil
}

type VSplitter struct{ BaseTile }

func (s *VSplitter) Interact(l *Light) (*Light, *Light) {
	if s.Seen(l) {
		return nil, nil
	}
	if l.Dir == West || l.Dir == East {
		l2 := *l
		l.Dir = North
		l2.Dir = South
		l.Move()
		l2.Move()
		return l, &l2
	}
	l.Move()
	return l, nil
}

type BackMirror struct{ BaseTile }

func (t *BackMirror) Interact(l *Light) (*Light, *Light) {
	if t.Seen(l) {
		return nil, nil
	}
	switch l.Dir {
	case North:
		l.Dir = West
	case East:
		l.Dir = South
	case South:
		l.Dir = East
	case West:
		l.Dir = North
	}
	l.Move()
	return l, nil
}

type ForwardMirror struct{ BaseTile }

func (t *ForwardMirror) Interact(l *Light) (*Light, *Light) {
	if t.Seen(l) {
		return nil, nil
	}
	switch l.Dir {
	case North:
		l.Dir = East
	case East:
		l.Dir = North
	case South:
		l.Dir = West
	case West:
		l.Dir = South
	}
	l.Move()
	return l, nil
}

type Board struct {
	rows, cols int

	tiles [][]Tile
}

func (b *Board) Clear() {
	for y := 0; y < b.rows; y++ {
		for x := 0; x < b.cols; x++ {
			b.tiles[y][x].Clear()
		}
	}
}

func load(fn string) *Board {
	var b Board
	data, err := os.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	lines := bytes.Split(data, []byte{'\n'})
	b.rows = len(lines)
	b.tiles = make([][]Tile, b.rows)
	for y, line := range lines {
		line = bytes.TrimSpace(line)
		if b.cols == 0 {
			b.cols = len(line)
		}
		row := make([]Tile, b.cols)
		for x, char := range line {
			switch char {
			case '|':
				row[x] = &VSplitter{}
			case '-':
				row[x] = &HSplitter{}
			case '\\':
				row[x] = &BackMirror{}
			case '/':
				row[x] = &ForwardMirror{}
			default:
				row[x] = &EmptyTile{}
			}
		}
		b.tiles[y] = row
	}
	return &b
}

func p1(board *Board) int {
	light := &Light{X: 0, Y: 0, Dir: East}
	return calc(board, light)
}

func calc(board *Board, light *Light) int {
	stack := []*Light{light}
	for len(stack) > 0 {
		l := stack[0]
		stack[0] = nil
		stack = stack[1:]
		l1, l2 := board.tiles[l.Y][l.X].Interact(l)
		if l1 != nil {
			if !(l1.X < 0 || l1.X >= board.cols || l1.Y < 0 || l1.Y >= board.rows) {
				stack = append(stack, l1)
			}
		}
		if l2 != nil {
			if !(l2.X < 0 || l2.X >= board.cols || l2.Y < 0 || l2.Y >= board.rows) {
				stack = append(stack, l2)
			}
		}
	}
	total := 0
	for y := 0; y < board.rows; y++ {
		for x := 0; x < board.cols; x++ {
			if board.tiles[y][x].IsEnergized() {
				total++
			}
		}
	}
	return total
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func p2(board *Board) int {
	total := 0
	for y := 0; y < board.rows; y++ {
		left := &Light{Y: y, X: 0, Dir: East}
		right := &Light{Y: y, X: board.cols - 1, Dir: West}
		board.Clear()
		total = max(total, calc(board, left))
		board.Clear()
		total = max(total, calc(board, right))
	}
	for x := 0; x < board.cols; x++ {
		top := &Light{Y: 0, X: x, Dir: South}
		bottom := &Light{Y: board.rows - 1, X: x, Dir: North}
		board.Clear()
		total = max(total, calc(board, top))
		board.Clear()
		total = max(total, calc(board, bottom))
	}
	return total
}
