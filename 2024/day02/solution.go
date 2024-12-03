package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := "assets/input.txt"
	p1(filename)
	p2(filename)
}

func p1(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	lines := bytes.Split(data, []byte{'\n'})
	var count int
	for _, line := range lines {
		var report Report
		if err := report.UnmarshalText(line); err != nil {
			panic(err)
		}
		if report.Safe(false) {
			count++
		}
	}
	fmt.Println(count)
}

func p2(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	lines := bytes.Split(data, []byte{'\n'})
	var count int
	for _, line := range lines {
		var report Report
		if err := report.UnmarshalText(line); err != nil {
			panic(err)
		}
		if report.Safe(true) {
			count++
		}
	}
	fmt.Println(count)
}

type Report struct {
	levels []int
}

func (r *Report) deltas() []int {
	out := make([]int, len(r.levels)-1)
	for i := 1; i < len(r.levels); i++ {
		out[i-1] = r.levels[i] - r.levels[i-1]
	}
	return out
}

func (r *Report) UnmarshalText(data []byte) error {
	fields := bytes.Fields(data)
	for _, field := range fields {
		v := mustParse(string(field))
		r.levels = append(r.levels, v)
	}
	return nil
}

func (r *Report) Safe(recursive bool) bool {
	if r._safe() {
		return true
	}
	if !recursive {
		return false
	}
	for i := range len(r.levels) {
		var _r Report
		_r.levels = append(_r.levels, r.levels[:i]...)
		_r.levels = append(_r.levels, r.levels[i+1:]...)
		if _r._safe() {
			return true
		}
	}
	return false
}

func (r *Report) _safe() bool {
	deltas := r.deltas()
	_, increasing := direction(deltas)
	if increasing {
		for _, delta := range deltas {
			if delta <= 0 || delta > 3 {
				return false
			}
		}
	} else {
		for _, delta := range deltas {
			if delta >= 0 || delta < -3 {
				return false
			}
		}
	}

	return true
}

func direction(deltas []int) (int, bool) {
	var numPos, numNeg, numSame int
	for _, delta := range deltas {
		switch {
		case delta == 0:
			numSame++
		case delta < 0:
			numNeg++
		case delta > 0:
			numPos++
		}
	}
	if numPos > numNeg {
		// if increasing return the number of non-increasing values
		return numNeg, true
	} else {
		// if decreasing, return the number of non-decreasing values
		return numPos, false
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func mustParse(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
