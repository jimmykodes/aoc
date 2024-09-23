package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines := load("assets/input.txt")
	fmt.Println(p1(lines))
	fmt.Println(p2(lines))
}

func p2(lines []*Line) int64 {
	var polygon []Coord
	var lineArea int64
	for _, line := range lines {
		polygon = append(polygon, line.ColorStart)
		lineArea += line.ColorLen
	}
	// I don't really understand how we get the off by one. I understand that the area
	// calculation doesn't take into account the actual permiter values, but I don't
	// really get why I have do divide the line area in half and add one.
	return calculatePolygonArea(polygon) + (lineArea / 2) + 1
}

func p1(lines []*Line) int {
	var polygon []Coord
	var maxX, minX, maxY, minY int
	var lineArea int
	for _, line := range lines {
		start := line.Start
		maxX = max(maxX, start.X)
		minX = min(minX, start.X)
		maxY = max(maxY, start.Y)
		minY = min(minY, start.Y)
		polygon = append(polygon, start)
		lineArea += line.Len
	}
	ic := insideCount(minY, maxY, minX, maxX, polygon)
	cpa := calculatePolygonArea(polygon)
	fmt.Println(ic, cpa+int64(lineArea/2)+1)
	return ic
}

func load(fn string) []*Line {
	data, err := os.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	dataLines := bytes.Split(data, []byte{'\n'})
	lines := make([]*Line, len(dataLines))
	for i, dataLine := range dataLines {
		var line Line
		line.UnmarshalText(dataLine)
		lines[i] = &line
	}
	for i, line := range lines {
		if i == 0 {
			line.Start = Coord{0, 0}
			line.ColorStart = Coord{0, 0}
		} else {
			line.Start = lines[i-1].End
			line.ColorStart = lines[i-1].ColorEnd
		}
		line.End = line.Start.Offset(line.Dir, line.Len)
		line.ColorEnd = line.ColorStart.Offset(line.ColorDir, int(line.ColorLen))
	}
	return lines
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func calculatePolygonArea(coords []Coord) int64 {
	n := len(coords)
	if n < 3 {
		// Invalid input
		return 0
	}

	var area int64

	for i := 0; i < n-1; i++ {
		area += int64(coords[i].X*coords[i+1].Y - coords[i+1].X*coords[i].Y)
	}

	// Add the last term
	area += int64(coords[n-1].X*coords[0].Y - coords[0].X*coords[n-1].Y)

	// Take the absolute value and divide by 2
	if area < 0 {
		area = -area
	}

	return area / 2
}

func insideCount(minY, maxY, minX, maxX int, polygon []Coord) int {
	total := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			point := Coord{X: x, Y: y}
			inside := false
			for i, j := 0, len(polygon)-1; i < len(polygon); j, i = i, i+1 {
				start := polygon[j]
				end := polygon[i]

				if start.X == end.X && point.X == start.X {
					if min(start.Y, end.Y) <= point.Y && point.Y <= max(start.Y, end.Y) {
						inside = true
						break
					}
				}

				if start.Y == end.Y && point.Y == start.Y {
					if min(start.X, end.X) <= point.X && point.X <= max(start.X, end.X) {
						inside = true
						break
					}
				}

				if (start.Y > point.Y) != (end.Y > point.Y) && point.X < (end.X-start.X)*(point.Y-start.Y)/(end.Y-start.Y)+start.X {
					inside = !inside
				}
			}
			if inside {
				total++
			}
		}
	}
	return total
}

type Direction int

const (
	Right Direction = iota + 1
	Down
	Left
	Up
)

var dirs = [...]Direction{
	'R': Right,
	'L': Left,
	'D': Down,
	'U': Up,
	'0': Right,
	'1': Down,
	'2': Left,
	'3': Up,
}

type Coord struct {
	Y, X int
}

func (c Coord) Offset(dir Direction, dist int) Coord {
	switch dir {
	case Right:
		c.X += dist
	case Down:
		c.Y += dist
	case Left:
		c.X -= dist
	case Up:
		c.Y -= dist
	}
	return c
}

type Line struct {
	Dir   Direction
	Len   int
	Start Coord
	End   Coord

	ColorDir   Direction
	ColorLen   int64
	ColorStart Coord
	ColorEnd   Coord
}

func (l *Line) UnmarshalText(b []byte) error {
	fields := bytes.Fields(b)
	l.Dir = dirs[fields[0][0]]
	l.Len, _ = strconv.Atoi(string(fields[1]))
	l.ColorDir = dirs[fields[2][7]]
	l.ColorLen, _ = strconv.ParseInt(string(fields[2][2:7]), 16, 64)
	return nil
}
