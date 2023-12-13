package main

import (
	"fmt"
	"os"
)

func main() {
	t := newTiles("assets/input.txt")
	fmt.Println(p2(t.Next))
}

func p1(next func() *Tile) float64 {
	var start *Tile
	board := [][]*Tile{make([]*Tile, 0)}
	for t := next(); t != nil; t = next() {
		if t.Y >= len(board) && t.Y > 0 {
			board = append(board, make([]*Tile, 0, len(board[t.Y-1])))
		}
		board[t.Y] = append(board[t.Y], t)
		if t.X > 0 {
			// look left
			if left := board[t.Y][t.X-1]; left.Allowed[East] && t.Allowed[West] {
				// if the left tile is allowed to connect east
				// and the current tile is either start (meaning we don't know)
				// or is also allowed to connect west, link the tiles
				left.Tiles[East] = t
				t.Tiles[West] = left
			}
		}
		if t.Y > 0 {
			// look up
			if up := board[t.Y-1][t.X]; up.Allowed[South] && t.Allowed[North] {
				up.Tiles[South] = t
				t.Tiles[North] = up
			}
		}

		if t.Start {
			start = t
		}
	}
	current := start
	var n *Tile
	var spaces int
	for {
		current.Visited = true
		for dir := North; dir <= West; dir++ {
			if n = current.Tiles[dir]; n != nil && !n.Visited {
				break
			}
		}
		if n == nil || n.Visited {
			return float64(spaces+1) / 2
		}
		current = n
		spaces++
	}
}

func p2(next func() *Tile) int {
	var start *Tile
	board := [][]*Tile{make([]*Tile, 0)}
	for t := next(); t != nil; t = next() {
		if t.Y >= len(board) && t.Y > 0 {
			board = append(board, make([]*Tile, 0, len(board[t.Y-1])))
		}
		board[t.Y] = append(board[t.Y], t)
		if t.X > 0 {
			// look left
			if left := board[t.Y][t.X-1]; left.Allowed[East] && t.Allowed[West] {
				// if the left tile is allowed to connect east
				// and the current tile is either start (meaning we don't know)
				// or is also allowed to connect west, link the tiles
				left.Tiles[East] = t
				t.Tiles[West] = left
			}
		}
		if t.Y > 0 {
			// look up
			if up := board[t.Y-1][t.X]; up.Allowed[South] && t.Allowed[North] {
				up.Tiles[South] = t
				t.Tiles[North] = up
			}
		}

		if t.Start {
			start = t
		}
	}
	current := start
	var n *Tile
	polygon := []*Tile{start}
	size := 0
	minx := start.X
	maxx := start.X
	miny := start.Y
	maxy := start.Y
	for {
		current.Visited = true
		for dir := North; dir <= West; dir++ {
			if n = current.Tiles[dir]; n != nil && !n.Visited {
				break
			}
		}
		if n == nil || n.Visited {
			break
		}
		n.Poly = true
		size++
		if n.Vertex {
			polygon = append(polygon, n)
			if n.X < minx {
				minx = n.X
			}
			if n.X > maxx {
				maxx = n.X
			}
			if n.Y < miny {
				miny = n.Y
			}
			if n.Y > maxy {
				maxy = n.Y
			}
		}
		current = n
	}

	total := 0
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			t := board[y][x]
			if t.Poly || t.Start {
				continue
			}
			if t.X < minx || t.X > maxx || t.Y < miny || t.Y > maxy {
				continue
			}
			inside := false
			for i, j := 0, len(polygon)-1; i < len(polygon); j, i = i, i+1 {
				if (polygon[i].Y > t.Y) != (polygon[j].Y > t.Y) &&
					t.X < (polygon[j].X-polygon[i].X)*(t.Y-polygon[i].Y)/(polygon[j].Y-polygon[i].Y)+polygon[i].X {
					inside = !inside
				}
			}
			if inside {
				total++
			}
		}
	}
	return total
}

const (
	North = iota
	East
	South
	West
)

type Tile struct {
	Start   bool
	X, Y    int
	Vertex  bool
	Visited bool
	Poly    bool
	Tiles   [4]*Tile
	Allowed [4]bool
}

func (t *Tile) Allow(dirs ...int) {
	for _, dir := range dirs {
		t.Allowed[dir] = true
	}
}

type Tiles struct {
	data    []byte
	pointer int
	x       int
	y       int
}

func newTiles(fn string) *Tiles {
	data, err := os.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	return &Tiles{data: data}
}

func (l *Tiles) Next() *Tile {
	if l.pointer >= len(l.data) {
		return nil
	}
	t := Tile{
		X: l.x,
		Y: l.y,
	}
	switch l.data[l.pointer] {
	case '-':
		t.Allow(East, West)
	case '|':
		t.Allow(North, South)
	case 'J':
		t.Allow(North, West)
		t.Vertex = true
	case 'L':
		t.Allow(North, East)
		t.Vertex = true
	case 'F':
		t.Allow(East, South)
		t.Vertex = true
	case '7':
		t.Allow(South, West)
		t.Vertex = true
	case 'S':
		t.Start = true
		t.Vertex = true
		t.Allow(North, East, South, West)
	case '\n':
		l.pointer++
		l.x = 0
		l.y++
		return l.Next()
	}
	l.pointer++
	l.x++
	return &t
}
