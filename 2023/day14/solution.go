package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	board := load("assets/input.txt")
	fmt.Println(p2(board))
}

func p1(board *Board) int {
	board.Tilt(0, -1)
	total := 0
	for _, rock := range board.Rocks {
		if !rock.Round {
			continue
		}
		total += board.Rows - rock.Y
	}
	return total
}

func p2(board *Board) int {
	cycles := 1000000000
	hashes := make(map[string]int)
	period_end := 0
	period_start := 0
	for i := 0; i < cycles; i++ {
		board.Tilt(0, -1)
		board.Tilt(-1, 0)
		board.Tilt(0, 1)
		board.Tilt(1, 0)
		hash := board.Hash()
		if cycle, ok := hashes[hash]; ok {
			period_start = cycle
			period_end = i
			break
		}
		hashes[hash] = i
	}
	period_len := period_end - period_start
	remaining := (cycles - period_start) - 1
	mod := remaining % period_len
	for i := 0; i < mod; i++ {
		board.Tilt(0, -1)
		board.Tilt(-1, 0)
		board.Tilt(0, 1)
		board.Tilt(1, 0)
	}
	total := 0
	for _, rock := range board.Rocks {
		if !rock.Round {
			continue
		}
		total += board.Rows - rock.Y
	}
	return total
}

type Rock struct {
	Round bool
	X, Y  int
	Step  int
	Board *Board
}

func (r *Rock) Move(dx, dy int) bool {
	r.Step = r.Board.Step
	if !r.Round {
		return false
	}
	x := r.X + dx
	y := r.Y + dy
	if x < 0 || x >= r.Board.Cols || y < 0 || y >= r.Board.Rows {
		return false
	}
	tile := r.Board.Tiles[y][x]
	if tile == nil {
		r.Board.Tiles[r.Y][r.X] = nil
		r.X = x
		r.Y = y
		r.Board.Tiles[y][x] = r
		return true
	}
	if !tile.Round {
		return false
	}
	if tile.Step == r.Step {
		// tile already moved this step, do nothing
		return false
	}
	tileMoved := tile.Move(dx, dy)
	thisMoved := r.Move(dx, dy)

	return tileMoved || thisMoved
}

type Board struct {
	Cols, Rows int
	Rocks      []*Rock
	Tiles      [][]*Rock
	Step       int
}

func (b *Board) Init() {
	b.Tiles = make([][]*Rock, b.Rows)
	for y := 0; y < b.Rows; y++ {
		b.Tiles[y] = make([]*Rock, b.Cols)
	}
	for _, r := range b.Rocks {
		b.Tiles[r.Y][r.X] = r
	}
}

func (b *Board) Tilt(dx, dy int) {
	moved := true
	for moved {
		b.Step++
		moved = false
		for _, rock := range b.Rocks {
			moved = moved || rock.Move(dx, dy)
		}
	}
}

func (b Board) Hash() string {
	// hash := sha1.New()
	// for _, rock := range b.Rocks {
	// 	hash.Write([]byte{byte(rock.X), byte(rock.Y)})
	// }
	// return hex.EncodeToString((hash.Sum(nil))[:])
	hash := sha1.Sum([]byte(b.String()))
	return hex.EncodeToString(hash[:])
}

func (b Board) String() string {
	var sb strings.Builder
	for _, row := range b.Tiles {
		for _, col := range row {
			if col == nil {
				sb.WriteRune('.')
			} else if col.Round {
				sb.WriteRune('O')
			} else {
				sb.WriteRune('#')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func load(fn string) *Board {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var board Board
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for col, char := range scanner.Bytes() {
			if board.Rows == 0 && col > board.Cols {
				board.Cols = col
			}
			switch char {
			case 'O':
				// round
				board.Rocks = append(board.Rocks, &Rock{Round: true, X: col, Y: board.Rows, Board: &board})
			case '#':
				// cube
				board.Rocks = append(board.Rocks, &Rock{Round: false, X: col, Y: board.Rows, Board: &board})
			}
		}
		board.Rows++
	}
	board.Cols++
	board.Init()
	return &board
}
