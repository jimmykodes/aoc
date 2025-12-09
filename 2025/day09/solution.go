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

var fname = "assets/test.txt"

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
	var polygon [][2]*Tile
	for i := 1; i < len(tiles); i++ {
		polygon = append(polygon, [2]*Tile{tiles[i-1], tiles[i]})
	}

	// wrap
	polygon = append(polygon, [2]*Tile{tiles[len(tiles)-1], tiles[0]})

	squares := make([]Square, 0, len(tiles)*len(tiles))
	for i, tile := range tiles {
		for _, otherTile := range tiles[i+1:] {
			if tile == otherTile {
				continue
			}
			squares = append(squares, Square{tile, otherTile})
		}
	}

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
SQUARES:
	for _, square := range squares {
		for _, line := range square.Lines() {
			for _, boundary := range polygon {
				if intersect(line, boundary) {
					// if we intersect any polygon line, we can't be inside it.
					continue SQUARES
				}
			}
		}
		fmt.Printf("square: %v does not intersect polygon\n", square)

		// not intersecting a line could still mean we are full outside the polygon
		// rather than fully inside it.
		midpoint := square.Midpoint()
		rays := [][2]*Tile{
			{&Tile{0, midpoint.Y}, midpoint},
			{midpoint, &Tile{math.MaxInt, midpoint.Y}},
			{&Tile{midpoint.X, 0}, midpoint},
			{midpoint, &Tile{midpoint.X, math.MaxInt}},
		}
		for _, ray := range rays {
			count := 0
			for _, boundary := range polygon {
				if intersect(ray, boundary) {
					count++
				}
			}
			if count%2 == 0 {
				fmt.Printf("square: %v: ray %v count %v\n", square, ray, count)
				continue SQUARES
			}
		}
		// if we tested every ray, and they were all odd intersections,
		// we are fully inside the polygon... right?
		fmt.Printf("%v square.Area(): %v\n", square, square.Area())
		for _, ray := range rays {
			count := 0
			for _, boundary := range polygon {
				if intersect(ray, boundary) {
					count++
				}
			}
			fmt.Printf("square: %v: ray %v count %v\n", square, ray, count)
		}
		return
	}
}

func intersect(line1, line2 [2]*Tile) bool {
	p1, q1 := line1[0], line1[1]
	p2, q2 := line2[0], line2[1]

	// Calculate direction vectors
	d1 := Tile{X: q1.X - p1.X, Y: q1.Y - p1.Y}
	d2 := Tile{X: q2.X - p2.X, Y: q2.Y - p2.Y}

	// Calculate determinant (cross product of direction vectors)
	det := d1.X*d2.Y - d1.Y*d2.X

	// Lines are parallel if determinant is 0
	if det == 0 {
		return false
	}

	// Calculate difference vector from p1 to p2
	dp := Tile{X: p2.X - p1.X, Y: p2.Y - p1.Y}

	// Calculate parameters t and u using Cramer's rule
	// t is parameter for line1: p1 + t*d1
	// u is parameter for line2: p2 + u*d2
	t := float64(dp.X*d2.Y-dp.Y*d2.X) / float64(det)
	u := float64(dp.X*d1.Y-dp.Y*d1.X) / float64(det)

	// Lines intersect if both parameters are strictly between 0 and 1
	// Using strict inequality excludes endpoint intersections (T-junctions)
	return t > 0 && t < 1 && u > 0 && u < 1
}

func _intersect(line1, line2 [2]*Tile) bool {
	p1, q1 := line1[0], line1[1]
	p2, q2 := line2[0], line2[1]

	// Helper function to find orientation of ordered triplet (p, q, r)
	// Returns:
	// 0 -> p, q and r are colinear
	// 1 -> Clockwise
	// 2 -> Counterclockwise
	orientation := func(p, q, r *Tile) int {
		val := (q.Y-p.Y)*(r.X-q.X) - (q.X-p.X)*(r.Y-q.Y)
		if val == 0 {
			return 0 // colinear
		}
		if val > 0 {
			return 1 // clockwise
		}
		return 2 // counterclockwise
	}

	// Check if point q lies on segment pr
	onSegment := func(p, q, r *Tile) bool {
		return q.X <= max(p.X, r.X) && q.X >= min(p.X, r.X) &&
			q.Y <= max(p.Y, r.Y) && q.Y >= min(p.Y, r.Y)
	}

	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)

	// Special cases - we return false for these since they don't count as "crossing"
	// according to your requirements (no T intersections or full overlaps)

	// p1, q1 and p2 are colinear and p2 lies on segment p1q1
	if o1 == 0 && onSegment(p1, p2, q1) {
		return false
	}

	// p1, q1 and q2 are colinear and q2 lies on segment p1q1
	if o2 == 0 && onSegment(p1, q2, q1) {
		return false
	}

	// p2, q2 and p1 are colinear and p1 lies on segment p2q2
	if o3 == 0 && onSegment(p2, p1, q2) {
		return false
	}

	// p2, q2 and q1 are colinear and q1 lies on segment p2q2
	if o4 == 0 && onSegment(p2, q1, q2) {
		return false
	}

	// General case - proper intersection
	if o1 != o2 && o3 != o4 {
		return true
	}

	return false
}

type Square [2]*Tile

func (e Square) Area() int {
	xDist := e[0].X - e[1].X
	if xDist < 0 {
		xDist = -xDist
	}

	yDist := e[0].Y - e[1].Y
	if yDist < 0 {
		yDist = -yDist
	}
	return (xDist + 1) * (yDist + 1)
}

func (e Square) Lines() [4][2]*Tile {
	p1 := e[0]
	p3 := e[1]
	p2 := &Tile{p1.X, p3.Y}
	p4 := &Tile{p3.X, p1.Y}

	return [4][2]*Tile{
		{p1, p2},
		{p2, p3},
		{p3, p4},
		{p4, p1},
	}
}

func (e Square) Midpoint() *Tile {
	x := (e[0].X + e[1].X) / 2
	y := (e[0].Y + e[1].Y) / 2
	return &Tile{x, y}
}

func (e Square) String() string {
	return fmt.Sprintf("%s -> %s", e[0], e[1])
}

type Tile struct {
	X, Y int
}

func (b Tile) String() string {
	return fmt.Sprintf("( %d, %d )", b.X, b.Y)
}

func (b *Tile) UnmarshalText(data []byte) error {
	x, y, _ := bytes.Cut(data, []byte{','})

	X, xErr := strconv.Atoi(string(x))
	Y, yErr := strconv.Atoi(string(y))

	b.X, b.Y = X, Y

	return errors.Join(xErr, yErr)
}

func parse(fname string) []*Tile {
	data, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(bytes.TrimSpace(data), []byte{'\n'})
	out := make([]*Tile, len(lines))
	for i, line := range lines {
		var b Tile
		if err := b.UnmarshalText(line); err != nil {
			panic(err)
		}
		out[i] = &b
	}
	return out
}
