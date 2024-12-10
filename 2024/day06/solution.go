package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	b, p := parseBoard("assets/test.txt")
	// p1(b, p)
	p2(b, p)
}

type state int

const (
	open state = iota
	visited
	obstructed
)

type dir int

const (
	up dir = iota
	right
	down
	left
	mod
)

type patrol struct {
	pos [2]int
	dir dir
}

func (p *patrol) Next(b board) *tile {
	x, y := p.pos[0], p.pos[1]
	switch p.dir {
	case up:
		if y == 0 {
			return nil
		}
		return b[y-1][x]
	case right:
		if x == len(b[0])-1 {
			return nil
		}
		return b[y][x+1]
	case down:
		if y == len(b)-1 {
			return nil
		}
		return b[y+1][x]
	case left:
		if x == 0 {
			return nil
		}
		return b[y][x-1]
	default:
		panic("invalid dir")
	}
}

func (p *patrol) Move(b board) bool {
	x, y := p.pos[0], p.pos[1]
	b.Visit(x, y, p.dir)
	if n := p.Next(b); n == nil {
		return false
	} else if n.state == obstructed {
		// rotate
		p.dir = p.turn()
		return p.Move(b)
	} else {
		// set position
		p.pos[0] = n.pos[0]
		p.pos[1] = n.pos[1]
		return true
	}
}

func (p *patrol) turn() dir {
	return (p.dir + 1) % mod
}

type tile struct {
	pos         [2]int
	state       state
	visitedDirs [mod]bool
	wouldLoop   bool
}
type board [][]*tile

func (b board) Visit(x, y int, d dir) {
	b[y][x].state = visited
	b[y][x].visitedDirs[d] = true
}

func parseBoard(file string) (board, patrol) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	lines := bytes.Split(data, []byte{'\n'})
	b := make(board, len(lines))
	var p patrol
	for row, line := range lines {
		b[row] = make([]*tile, len(line))
		for col, val := range line {
			b[row][col] = &tile{pos: [2]int{col, row}}
			switch val {
			case '.':
				b[row][col].state = open
			case '#':
				b[row][col].state = obstructed
			case '>':
				p.pos[1] = row
				p.pos[0] = col
				p.dir = right
			case '^':
				p.pos[1] = row
				p.pos[0] = col
				p.dir = up
			case '<':
				p.pos[1] = row
				p.pos[0] = col
				p.dir = left
			case 'v':
				p.pos[1] = row
				p.pos[0] = col
				p.dir = down
			}
		}
	}

	return b, p
}

func p1(b board, p patrol) {
	pat := &p
	for pat.Move(b) {
	}
	total := 0
	for _, row := range b {
		for _, col := range row {
			if col.state == visited {
				total++
			}
		}
	}
	fmt.Println(total)
}

func p2(b board, p patrol) {
	pat := &p
	n := pat.Next(b)
	for n != nil {
		if n.visitedDirs[p.turn()] {
			// turning on the next tile would create a loop
			x, y := n.pos[0], n.pos[1]
			switch p.dir {
			case up:
				b[y-1][x].wouldLoop = true
			case right:
				b[y][x+1].wouldLoop = true
			case down:
				b[y+1][x].wouldLoop = true
			case left:
				b[y][x-1].wouldLoop = true
			}
		}
		pat.Move(b)
		n = pat.Next(b)
	}
	total := 0
	for _, row := range b {
		for _, col := range row {
			if col.wouldLoop {
				total++
			}
		}
	}
	fmt.Println(total)
}
