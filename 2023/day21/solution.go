package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	start := load("assets/input.txt")
	fmt.Println(p1(start, 64))
}

func p1(start *Tile, steps int) int {
	stack := map[int]*Tile{start.ID: start}
	for i := 0; i < steps; i++ {
		newStack := make(map[int]*Tile, len(stack)*4)
		for _, tile := range stack {
			if tile.North != nil {
				newStack[tile.North.ID] = tile.North
			}
			if tile.East != nil {
				newStack[tile.East.ID] = tile.East
			}
			if tile.West != nil {
				newStack[tile.West.ID] = tile.West
			}
			if tile.South != nil {
				newStack[tile.South.ID] = tile.South
			}
		}
		stack = newStack
	}
	return len(stack)
}

type Coord struct {
	Y, X int
}
type Tile struct {
	ID    int
	Coord Coord

	North, East, South, West *Tile
}

func load(fn string) *Tile {
	data, err := os.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	lines := bytes.Split(data, []byte{'\n'})
	tiles := make([][]*Tile, len(lines))
	for i := range tiles {
		tiles[i] = make([]*Tile, len(lines[0]))
	}
	var start *Tile
	var id int
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				// rock, do nothing
				continue
			}
			t := &Tile{ID: id, Coord: Coord{Y: y, X: x}}
			id++
			if y > 0 {
				// look north
				if north := tiles[y-1][x]; north != nil {
					t.North = north
					north.South = t
				}
			}
			if x > 0 {
				// look west
				if west := tiles[y][x-1]; west != nil {
					t.West = west
					west.East = t
				}
			}
			if char == 'S' {
				start = t
			}
			tiles[y][x] = t
		}
	}
	return start
}
