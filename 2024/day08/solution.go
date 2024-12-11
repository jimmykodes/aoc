package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	antennai, rows, cols := puzzle("assets/input.txt")
	p2(antennai, rows, cols)
}

func p2(antennai map[byte][]*Antenna, rows int, cols int) {
	antiNodes := map[byte][]*Antenna{}
	for freq, ant := range antennai {
		for i := range len(ant) {
			a := ant[i]
			for _, other := range ant {
				if other == a {
					continue
				}
				dx := a.Run(other)
				dy := a.Rise(other)
				last := a
				for {
					antinodeA := &Antenna{X: last.X + dx, Y: last.Y + dy, Frequency: freq}
					if !antinodeA.InBounds(rows, cols) {
						break
					}
					antiNodes[freq] = append(antiNodes[freq], antinodeA)
					last = antinodeA

				}
			}
		}
	}
	b := make([][]struct {
		Freq byte
		Anti byte
	}, rows)
	for col := range cols {
		b[col] = make([]struct {
			Freq byte
			Anti byte
		}, cols)
	}
	for freq, ant := range antennai {
		for _, a := range ant {
			b[a.Y][a.X].Freq = freq
		}
	}
	for freq, ant := range antiNodes {
		for _, a := range ant {
			b[a.Y][a.X].Anti = freq
		}
	}
	var total int
	for _, row := range b {
		for _, col := range row {
			if col.Freq == 0 && col.Anti == 0 {
				fmt.Print(".")
			} else if col.Freq != 0 {
				fmt.Print(string(col.Freq))
			} else if col.Anti != 0 {
				// fmt.Print(string(col.Anti))
				fmt.Print("#")
				total++
			}
		}
		fmt.Println()
	}
	for _, ant := range antennai {
		if l := len(ant); l > 1 {
			total += l
		}
	}
	fmt.Println(total)
}

func p1(antennai map[byte][]*Antenna, rows int, cols int) {
	antiNodes := map[byte][]*Antenna{}
	for freq, ant := range antennai {
		for i := range len(ant) {
			a := ant[i]
			for b := i + 1; b < len(ant); b++ {
				other := ant[b]
				dx := a.Run(other)
				dy := a.Rise(other)
				antinodeA := &Antenna{X: a.X + dx, Y: a.Y + dy, Frequency: freq}
				antinodeB := &Antenna{X: other.X - dx, Y: other.Y - dy, Frequency: freq}
				if antinodeA.InBounds(rows, cols) {
					antiNodes[freq] = append(antiNodes[freq], antinodeA)
				}
				if antinodeB.InBounds(rows, cols) {
					antiNodes[freq] = append(antiNodes[freq], antinodeB)
				}
			}
		}
	}
	b := make([][]struct {
		Freq byte
		Anti bool
	}, rows)
	for col := range cols {
		b[col] = make([]struct {
			Freq byte
			Anti bool
		}, cols)
	}
	for freq, ant := range antennai {
		for _, a := range ant {
			b[a.Y][a.X].Freq = freq
		}
	}
	for freq, ant := range antiNodes {
		for _, a := range ant {
			b[a.Y][a.X].Anti = true
			if b[a.Y][a.X].Freq == 0 {
				b[a.Y][a.X].Freq = freq
			}
		}
	}
	var total int
	for _, row := range b {
		for _, col := range row {
			if col.Freq == 0 {
				fmt.Print(".")
			} else {
				if col.Anti {
					total++
					fmt.Print("#")
				} else {
					fmt.Print(string(col.Freq))
				}
			}
		}
		fmt.Println()
	}
	fmt.Println(total)
}

type Antenna struct {
	Frequency byte
	X, Y      int
}

func (a Antenna) Rise(b *Antenna) int {
	return a.Y - b.Y
}

func (a Antenna) Run(b *Antenna) int {
	return a.X - b.X
}

func (a Antenna) InBounds(numRows, numCols int) bool {
	return 0 <= a.X && a.X < numCols && 0 <= a.Y && a.Y < numRows
}

func puzzle(filename string) (map[byte][]*Antenna, int, int) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	lines := bytes.Fields(data)
	numRows := len(lines)
	numCols := len(lines[0])
	antennai := make(map[byte][]*Antenna)
	for rowNum, line := range lines {
		for colNum, freq := range line {
			switch freq {
			case '.':
				continue
			default:
				antennai[freq] = append(antennai[freq], &Antenna{
					X:         colNum,
					Y:         rowNum,
					Frequency: freq,
				})
			}
		}
	}

	return antennai, numRows, numCols
}
