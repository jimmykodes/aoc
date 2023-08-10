package main

import (
	"fmt"
	"strconv"
	"strings"
)

func NewCoord(x, y int) *Coord {
	return &Coord{X: x, Y: y}
}

type Coord struct {
	X, Y int
}

func (c *Coord) String() string {
	return fmt.Sprintf("<Coord (%d, %d)>", c.X, c.Y)
}

func (c *Coord) Interpolate(other *Coord) []*Coord {
	switch {
	case c.X == other.X:
		// interpolate columns
		diff := c.Y - other.Y
		op := sub
		if diff < 0 {
			op = add
			diff *= -1
		}
		out := make([]*Coord, diff-1)
		for i := 0; i < diff-1; i++ {
			out[i] = NewCoord(c.X, op(c.Y, i+1))
		}
		return out

	case c.Y == other.Y:
		// interpolate row
		diff := c.X - other.X
		op := sub
		if diff < 0 {
			op = add
			diff *= -1
		}
		out := make([]*Coord, diff-1)
		for i := 0; i < diff-1; i++ {
			out[i] = NewCoord(op(c.X, i+1), c.Y)
		}
		return out
	default:
		panic("cannot interpolate diagonal lines")
	}
}

type Line struct {
	points []*Coord
}

func NewLineFromStr(line string) *Line {
	pointsStr := strings.Split(line, " -> ")
	var l Line
	for _, p := range pointsStr {
		xStr, yStr, _ := strings.Cut(p, ",")
		x, _ := strconv.Atoi(xStr)
		y, _ := strconv.Atoi(yStr)
		p := NewCoord(x, y)
		if len(l.points) > 0 {
			l.points = append(l.points, l.points[len(l.points)-1].Interpolate(p)...)
		}
		l.points = append(l.points, p)
	}
	return &l
}

func add(n, delta int) int {
	return n + delta
}
func sub(n, delta int) int {
	return n - delta
}
