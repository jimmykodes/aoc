package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	board := ParseFile("assets/input.txt")
	var p1 int

	for _, row := range board {
		for _, cell := range row {
			if cell.Accessible() {
				p1++
			}
		}
	}

	fmt.Println("p1", p1)
	var p2 int
	for {
		var accessible []*Space
		for _, row := range board {
			for _, cell := range row {
				if cell.Accessible() {
					accessible = append(accessible, cell)
				}
			}
		}
		if len(accessible) == 0 {
			break
		}
		p2 += len(accessible)
		for _, cell := range accessible {
			cell.Occupied = false
		}
	}
	fmt.Println("p2", p2)
}

type Space struct {
	Left, Right, Up, Down *Space
	Occupied              bool
}

const (
	upLeft = iota
	up
	upRight
	right
	downRight
	down
	downLeft
	left
)

func (s *Space) Neighbors() [8]*Space {
	out := [8]*Space{
		up:    s.Up,
		right: s.Right,
		down:  s.Down,
		left:  s.Left,
	}

	if s.Up != nil {
		out[upLeft] = s.Up.Left
		out[upRight] = s.Up.Right
	}

	if s.Down != nil {
		out[downRight] = s.Down.Right
		out[downLeft] = s.Down.Left
	}

	return out
}

func (s *Space) Accessible() bool {
	if !s.Occupied {
		return false
	}
	var count int
	for _, neighbor := range s.Neighbors() {
		if neighbor != nil && neighbor.Occupied {
			count++
		}
	}
	return count < 4
}

func ParseFile(name string) [][]*Space {
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(data, []byte{'\n'})

	out := make([][]*Space, len(lines))
	for y, line := range lines {
		row := make([]*Space, len(line))
		for x, val := range line {
			// space is considered occupied if there is a `@` symbol
			// anything else is unoccupied
			space := &Space{Occupied: val == '@'}

			if x > 0 {
				// we have something to our left,
				// make sure the left->right connections get made
				space.Left = row[x-1]
				row[x-1].Right = space
			}

			if y > 0 {
				// we have something above us,
				// make sure the up->down connections get made
				space.Up = out[y-1][x]
				out[y-1][x].Down = space
			}

			row[x] = space
		}
		out[y] = row
	}

	return out
}
