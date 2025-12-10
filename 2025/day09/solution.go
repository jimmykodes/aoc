package main

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

var fname = "assets/input.txt"

func main() {
	fmt.Println("p1")
	P1(fname)
	fmt.Println("p2")
	P2(fname)
}

func P1(fname string) {
	tiles := parse(fname)

	squares := make([]Square, 0, len(tiles)*len(tiles))
	for i, tile := range tiles {
		for _, otherTile := range tiles[i+1:] {
			if tile == otherTile {
				continue
			}
			squares = append(squares, Square{tile, otherTile})
		}
	}

	squares = slices.Clip(squares)

	slices.SortStableFunc(squares, func(a, b Square) int {
		aDist := a.Area()
		bDist := b.Area()
		switch {
		case aDist < bDist:
			return -1
		case aDist > bDist:
			return 1
		default:
			return 0
		}
	})

	largest := squares[len(squares)-1]
	fmt.Println(largest, largest.Area())
}

func P2(fname string) {
	tiles := parse(fname)
	var polygon []Line
	for i := 1; i < len(tiles); i++ {
		polygon = append(polygon, Line{tiles[i-1], tiles[i]})
	}

	// wrap
	polygon = append(polygon, [2]*Coord{tiles[len(tiles)-1], tiles[0]})

	squares := make([]Square, 0, len(tiles)*len(tiles))
	for i, tile := range tiles {
		for _, otherTile := range tiles[i+1:] {
			if tile == otherTile {
				continue
			}
			squares = append(squares, Square{tile, otherTile})
		}
	}

	squares = slices.Clip(squares)

	// squares := []Square{{&Coord{9, 5}, &Coord{2, 3}}}

	slices.SortStableFunc(squares, func(a, b Square) int {
		aDist := a.Area()
		bDist := b.Area()
		switch {
		case aDist < bDist:
			return 1
		case aDist > bDist:
			return -1
		default:
			return 0
		}
	})

	for _, square := range squares {
		if square.InPolygon(polygon) {
			fmt.Println(square, square.Area())
			return
		}
	}
}

type (
	Line   [2]*Coord
	Square [2]*Coord
)

func (s Square) InPolygon(polygon []Line) bool {
	for _, point := range s.Points() {
		if !point.InPolygon(polygon) {
			return false
		}
	}
	for _, line := range s.Lines() {
		for _, boundary := range polygon {
			if intersection := line.Intersection(boundary); intersection != nil {
				// lines are orthoganal
				if line.PointOnSegment(intersection) && boundary.PointOnSegment(intersection) {
					// lines cross
					return false
				}
				// lines don't cross, look at next boundary
				continue
			} else {
				// lines are parallel
				if line.FullyContains(boundary) {
					// lines are colinear
					if line[0].Eq(boundary[0]) {
						// start points match. check next boundary
						continue
					}
					if line[1].Eq(boundary[0]) {
						// line end matches boundary start
						// due to directionality, we can say
						// this means the line stays "inside"
						continue
					}
					return false
				} else {
					// lines are independent, look at next boundary
					continue
				}
			}
		}
	}
	return true
}

func (l Line) IsVertical() bool {
	return l[0].X == l[1].X
}

func (l Line) IsHorizontal() bool {
	return l[0].Y == l[1].Y
}

func (l Line) Intersection(other Line) *Coord {
	if l.IsHorizontal() && other.IsHorizontal() || l.IsVertical() && other.IsVertical() {
		// lines are parallel, no Intersection
		// or, technically, complete intersection if they
		// overlap, but that has no real bearing here
		return nil
	}
	if l.IsHorizontal() {
		return &Coord{other[0].X, l[0].Y}
	}
	return &Coord{l[0].X, other[0].Y}
}

func (l Line) maxX() int {
	return max(l[0].X, l[1].X)
}

func (l Line) maxY() int {
	return max(l[0].Y, l[1].Y)
}

func (l Line) minX() int {
	return min(l[0].X, l[1].X)
}

func (l Line) minY() int {
	return min(l[0].Y, l[1].Y)
}

func (l Line) FullyContains(other Line) bool {
	if l.IsHorizontal() && other.IsHorizontal() {
		return l[0].Y == other[0].Y && l.minX() <= other.minX() && other.maxX() <= l.maxX()
	}
	return l[0].X == other[0].X && l.minY() <= other.minY() && other.maxY() <= l.maxY()
}

func (l Line) PointOnSegment(c *Coord) bool {
	if c == nil {
		return false
	}
	if l.IsHorizontal() {
		return min(l[0].X, l[1].X) < c.X &&
			c.X < max(l[0].X, l[1].X)
	}
	return min(l[0].Y, l[1].Y) < c.Y &&
		c.Y < max(l[0].Y, l[1].Y)
}

func (s Square) Points() [4]*Coord {
	p1 := s[0]
	p3 := s[1]
	p2 := &Coord{p3.X, p1.Y}
	p4 := &Coord{p1.X, p3.Y}

	return [4]*Coord{p1, p2, p3, p4}
}

func (s Square) Area() int {
	xDist := s[0].X - s[1].X
	if xDist < 0 {
		xDist = -xDist
	}

	yDist := s[0].Y - s[1].Y
	if yDist < 0 {
		yDist = -yDist
	}
	return (xDist + 1) * (yDist + 1)
}

func (s Square) Lines() [4]Line {
	coords := s.Points()
	p1 := coords[0]
	p2 := coords[1]
	p3 := coords[2]
	p4 := coords[3]

	return [4]Line{
		{p1, p2},
		{p2, p3},
		{p3, p4},
		{p4, p1},
	}
}

func (s Square) Midpoint() *Coord {
	x := (s[0].X + s[1].X) / 2
	y := (s[0].Y + s[1].Y) / 2
	return &Coord{x, y}
}

func (s Square) String() string {
	return fmt.Sprintf("%s -> %s", s[0], s[1])
}

type Coord struct {
	X, Y int
}

func (c *Coord) InPolygon(polygon []Line) bool {
	rays := []Line{
		{&Coord{0, c.Y}, c},
		{c, &Coord{math.MaxInt, c.Y}},
		{&Coord{c.X, 0}, c},
		{c, &Coord{c.X, math.MaxInt}},
	}
	for _, ray := range rays {
		count := 0
		for _, boundary := range polygon {
			if c.Eq(boundary[0]) || c.Eq(boundary[1]) {
				// a point is in the polygon if it sits at a boundary vertex
				return true
			}
			intersection := ray.Intersection(boundary)
			if boundary.PointOnSegment(intersection) && ray.PointOnSegment(intersection) {
				count++
			} else if intersection == nil && ray.FullyContains(boundary) {
				count++
			}
		}
		if count%2 == 0 {
			return false
		}
	}
	return true
}

func (c Coord) Eq(other *Coord) bool {
	return c.X == other.X && c.Y == other.Y
}

func (c Coord) String() string {
	return fmt.Sprintf("( %d, %d )", c.X, c.Y)
}

func (c *Coord) UnmarshalText(data []byte) error {
	x, y, _ := bytes.Cut(data, []byte{','})

	X, xErr := strconv.Atoi(string(x))
	Y, yErr := strconv.Atoi(string(y))

	c.X, c.Y = X, Y

	return errors.Join(xErr, yErr)
}

func parse(fname string) []*Coord {
	data, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(bytes.TrimSpace(data), []byte{'\n'})
	out := make([]*Coord, len(lines))
	for i, line := range lines {
		var b Coord
		if err := b.UnmarshalText(line); err != nil {
			panic(err)
		}
		out[i] = &b
	}
	return out
}
