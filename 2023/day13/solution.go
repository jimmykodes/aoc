package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pats := load("assets/input.txt")
	fmt.Println(p1(pats))
	fmt.Println(p2(pats))
}

func p1(patterns []Pattern) int {
	var rtotal, vtotal int

PATTERN:
	for _, pattern := range patterns {
		for i := 1; i < len(pattern.rows); i++ {
			if pattern.Reflect(i-1, i, len(pattern.rows), pattern.Row) {
				rtotal += i
				continue PATTERN
			}
		}
		for i := 1; i < len(pattern.rows[0]); i++ {
			if pattern.Reflect(i-1, i, len(pattern.rows[0]), pattern.Col) {
				vtotal += i
				continue PATTERN
			}
		}
	}
	return vtotal + (100 * rtotal)
}

func p2(patterns []Pattern) int {
	var rtotal, vtotal int

PATTERN:
	for _, pattern := range patterns {
		for i := 1; i < len(pattern.rows[0]); i++ {
			if pattern.SmudgedReflect(i-1, i, len(pattern.rows[0]), pattern.Col) {
				vtotal += i
				continue PATTERN
			}
		}
		for i := 1; i < len(pattern.rows); i++ {
			if pattern.SmudgedReflect(i-1, i, len(pattern.rows), pattern.Row) {
				rtotal += i
				continue PATTERN
			}
		}
	}
	return vtotal + (100 * rtotal)
}

type Pattern struct {
	rows [][]bool
}

func (p Pattern) Row(idx int) uint32 {
	var out uint32
	for i, char := range p.rows[idx] {
		if char {
			out |= 1 << i
		}
	}

	return out
}

func (p Pattern) Col(idx int) uint32 {
	var out uint32
	for i, row := range p.rows {
		if row[idx] {
			out |= 1 << i
		}
	}
	return out
}

const (
	Vert = iota + 1
	Horiz
)

func (p Pattern) SmudgedReflect(a, b, max int, check func(int) uint32) bool {
	smudgeFixed := false
	for a >= 0 && b < max {
		tr := check(a)
		br := check(b)
		if tr != br {
			if dif := tr ^ br; dif&(dif-1) == 0 && !smudgeFixed {
				smudgeFixed = true
				a--
				b++
				continue
			}
			return false
		}
		a--
		b++
	}
	return true && smudgeFixed
}

func (p Pattern) Reflect(a, b, max int, check func(int) uint32) bool {
	for a >= 0 && b < max {
		tr := check(a)
		br := check(b)
		if tr != br {
			return false
		}
		a--
		b++
	}
	return true
}

func (p Pattern) XReflect(x1, x2 int) bool {
	top, bottom := x1, x2
	smudgedFixed := false
	for top >= 0 && bottom < len(p.rows[0]) {
		tr := p.Col(top)
		br := p.Col(bottom)
		if tr != br {
			if dif := tr &^ br; dif&(dif-1) == 0 && !smudgedFixed {
				smudgedFixed = true
				top--
				bottom++
				continue
			}
			return false
		}
		top--
		bottom++
	}
	return true && smudgedFixed
}

func (p Pattern) YReflect(y1, y2 int) bool {
	top, bottom := y1, y2
	smudgedFixed := false
	for top >= 0 && bottom < len(p.rows) {
		tr := p.Row(top)
		br := p.Row(bottom)
		if tr != br {
			if dif := tr &^ br; dif&(dif-1) == 0 && !smudgedFixed {
				smudgedFixed = true
				top--
				bottom++
				continue
			}
			return false
		}
		top--
		bottom++
	}
	return true && smudgedFixed
}

func (p Pattern) String() string {
	var rows [][]byte
	for _, row := range p.rows {
		r := make([]byte, len(row))
		for i, char := range row {
			if char {
				r[i] = '#'
			} else {
				r[i] = '.'
			}
		}
		rows = append(rows, r)
	}
	return string(bytes.Join(rows, []byte("\n")))
}

func (p Pattern) Bin() string {
	var rows []string

	for i := range p.rows {
		r := p.Row(i)
		rows = append(rows, strconv.FormatInt(int64(r), 2))
	}
	return strings.Join(rows, "\n")
}

func load(fn string) []Pattern {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var out []Pattern
	var pat Pattern
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			out = append(out, pat)
			pat = Pattern{}
			continue
		}
		row := make([]bool, len(line))
		for i, char := range line {
			row[i] = char == '#'
		}
		pat.rows = append(pat.rows, row)
	}
	out = append(out, pat)
	return out
}
