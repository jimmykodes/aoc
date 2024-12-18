package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	robots := puzzle("assets/input.txt")
	p2(robots)
}

func p1(robots []*Robot) {
	width, height := 101, 103
	for range 100 {
		for _, robot := range robots {
			robot.Update(width, height)
		}
	}
	out := make([][]int, height)
	for i := range height {
		out[i] = make([]int, width)
	}
	var q1, q2, q3, q4 int
	for _, robot := range robots {
		out[robot.Y][robot.X]++
		if robot.X < width/2 {
			if robot.Y < height/2 {
				q1++
			} else if robot.Y > height/2 {
				q2++
			}
		} else if robot.X > width/2 {
			if robot.Y < height/2 {
				q3++
			} else if robot.Y > height/2 {
				q4++
			}
		}
	}
	for y, row := range out {
		for x, col := range row {
			if x == width/2 || y == height/2 {
				fmt.Print(" ")
			} else if col == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
	fmt.Println(q1, q2, q3, q4)
	fmt.Println(q1 * q2 * q3 * q4)
}

func p2(robots []*Robot) {
	width, height := 101, 103
	i := 0
	for {
		i++
		for _, robot := range robots {
			robot.Update(width, height)
		}
		out := make([][]int, height)
		for i := range height {
			out[i] = make([]int, width)
		}

		for _, robot := range robots {
			out[robot.Y][robot.X]++
		}

		for _, row := range out[len(out)/2:] {
			if sum(row) > width/3 {
				for _, row := range out {
					for _, col := range row {
						if col == 0 {
							fmt.Print(".")
						} else {
							fmt.Print(col)
						}
					}
					fmt.Println()
				}
				fmt.Println()
				fmt.Println(i)
			}
		}

	}
}

func sum(ints []int) int {
	out := 0
	for _, i := range ints {
		out += i
	}
	return out
}

func puzzle(filename string) []*Robot {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	var out []*Robot
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		var r Robot
		if err := r.UnmarshalText(line); err != nil {
			panic(err)
		}
		out = append(out, &r)
	}
	return out
}

type Robot struct {
	X, Y    int
	DX, DY  int
	Checked bool
}

func (r *Robot) Update(width, height int) {
	r.Checked = false
	// move
	r.X += r.DX
	r.Y += r.DY

	// wrap
	if r.X >= width {
		r.X = r.X - width
	} else if r.X < 0 {
		r.X = width + r.X
	}
	if r.Y >= height {
		r.Y = r.Y - height
	} else if r.Y < 0 {
		r.Y = height + r.Y
	}
}

func (r Robot) String() string {
	return fmt.Sprintf("<Robot: pos %d,%d - vel %d,%d>", r.X, r.Y, r.DX, r.DY)
}

func (r *Robot) UnmarshalText(data []byte) error {
	pos, vel, _ := bytes.Cut(data, []byte{' '})
	x, y, _ := bytes.Cut(bytes.TrimPrefix(pos, []byte("p=")), []byte{','})
	var err error
	r.X, err = strconv.Atoi(string(x))
	if err != nil {
		return err
	}

	r.Y, err = strconv.Atoi(string(y))
	if err != nil {
		return err
	}

	dx, dy, _ := bytes.Cut(bytes.TrimPrefix(vel, []byte("v=")), []byte{','})
	r.DX, err = strconv.Atoi(string(dx))
	if err != nil {
		return err
	}

	r.DY, err = strconv.Atoi(string(dy))
	if err != nil {
		return err
	}
	return nil
}
