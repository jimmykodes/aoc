package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed test.txt
	test string
	//go:embed input.txt
	input string

	data = input
)

func main() {
	fmt.Println(p1())
	fmt.Println(p2())
}

func getLines() []*Line {
	lines := strings.Split(data, "\n")
	out := make([]*Line, len(lines))
	for i, line := range lines {
		out[i] = NewLineFromStr(line)
	}
	return out
}

func p1() int {
	lines := getLines()
	var maxX, maxY int
	for _, line := range lines {
		for _, point := range line.points {
			if point.X > maxX {
				maxX = point.X
			}
			if point.Y > maxY {
				maxY = point.Y
			}
		}
	}
	b := NewBoard(maxX, maxY)
	for _, line := range lines {
		for _, point := range line.points {
			b.board[point.Y][point.X] = Wall
		}
	}
	total := 0
	for {
		sand := &Coord{X: 500, Y: 0}
		for {
			// look down
			if sand.Y+1 >= len(b.board) {
				return total
			} else if b.board[sand.Y+1][sand.X] == Air {
				sand.Y += 1
			} else if b.board[sand.Y+1][sand.X-1] == Air {
				sand.Y += 1
				sand.X -= 1
			} else if b.board[sand.Y+1][sand.X+1] == Air {
				sand.Y += 1
				sand.X += 1
			} else {
				b.board[sand.Y][sand.X] = Sand
				total += 1
				break
			}
		}
	}
}
func p2() int {
	lines := getLines()
	var maxX, maxY int
	for _, line := range lines {
		for _, point := range line.points {
			if point.X > maxX {
				maxX = point.X
			}
			if point.Y > maxY {
				maxY = point.Y
			}
		}
	}
	b := NewBoard(maxX, maxY)
	for _, line := range lines {
		for _, point := range line.points {
			b.board[point.Y][point.X] = Wall
		}
	}
	b.board = append(b.board, make([]State, len(b.board[0])))
	for i := 0; i < len(b.board[0]); i++ {
		b.board[len(b.board)-1][i] = Wall
	}
	total := 0
	for {
		if b.board[0][500] == Sand {
			return total
		}
		sand := &Coord{X: 500, Y: 0}
		for {
			// look down
			if b.board[sand.Y+1][sand.X] == Air {
				sand.Y += 1
			} else if b.board[sand.Y+1][sand.X-1] == Air {
				sand.Y += 1
				sand.X -= 1
			} else if b.board[sand.Y+1][sand.X+1] == Air {
				sand.Y += 1
				sand.X += 1
			} else {
				b.board[sand.Y][sand.X] = Sand
				total += 1
				break
			}
		}
	}
}
