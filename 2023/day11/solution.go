package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	galaxies := load("assets/input.txt")
	// fmt.Println(p1(galaxies))
	fmt.Println(p2(galaxies))
}

func p2(galaxies []*Galaxy) int {
	var xs, ys []bool

	for _, gal := range galaxies {
		if dx := gal.Position.X - (len(xs) - 1); dx > 0 {
			dxs := make([]bool, dx)
			xs = append(xs, dxs...)
		}
		xs[gal.Position.X] = true

		if dy := gal.Position.Y - (len(ys) - 1); dy > 0 {
			dys := make([]bool, dy)
			ys = append(ys, dys...)
		}
		ys[gal.Position.Y] = true
	}

	for i := len(xs) - 1; i >= 0; i-- {
		if xs[i] {
			continue
		}
		for _, gal := range galaxies {
			if gal.Position.X > i {
				gal.Position.X += 999999
			}
		}
	}
	for i := len(ys) - 1; i >= 0; i-- {
		if ys[i] {
			continue
		}
		for _, gal := range galaxies {
			if gal.Position.Y > i {
				gal.Position.Y += 999999
			}
		}
	}

	total := 0

	for _, gal := range galaxies {
		for _, other := range galaxies {
			if other.ID <= gal.ID {
				continue
			}
			delta := gal.Position.Dist(other.Position)
			total += delta
		}
	}

	return total
}

func p1(galaxies []*Galaxy) int {
	var xs, ys []bool

	for _, gal := range galaxies {
		if dx := gal.Position.X - (len(xs) - 1); dx > 0 {
			dxs := make([]bool, dx)
			xs = append(xs, dxs...)
		}
		xs[gal.Position.X] = true

		if dy := gal.Position.Y - (len(ys) - 1); dy > 0 {
			dys := make([]bool, dy)
			ys = append(ys, dys...)
		}
		ys[gal.Position.Y] = true
	}

	for i := len(xs) - 1; i >= 0; i-- {
		if xs[i] {
			continue
		}
		for _, gal := range galaxies {
			if gal.Position.X > i {
				gal.Position.X++
			}
		}
	}
	for i := len(ys) - 1; i >= 0; i-- {
		if ys[i] {
			continue
		}
		for _, gal := range galaxies {
			if gal.Position.Y > i {
				gal.Position.Y++
			}
		}
	}

	total := 0

	for _, gal := range galaxies {
		for _, other := range galaxies {
			if other.ID <= gal.ID {
				continue
			}
			delta := gal.Position.Dist(other.Position)
			total += delta
		}
	}

	return total
}

type Coord struct {
	X, Y int
}

func (c Coord) Dist(other Coord) int {
	dx := other.X - c.X
	dy := other.Y - c.Y
	dx = int(math.Abs(float64(dx)))
	dy = int(math.Abs(float64(dy)))
	return dx + dy
}

type Galaxy struct {
	ID       int
	Position Coord
}

func load(fn string) []*Galaxy {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var out []*Galaxy
	scanner := bufio.NewScanner(f)
	y := 0
	id := 1
	for scanner.Scan() {
		for x, char := range scanner.Text() {
			if char == '#' {
				out = append(out, &Galaxy{ID: id, Position: Coord{X: x, Y: y}})
				id++
			}
		}
		y++
	}
	return out
}
