package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	p := puzzle("assets/input.txt")
	p1(p)
}

func puzzle(filename string) []int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var out []int
	for _, stone := range bytes.Fields(data) {
		v, err := strconv.Atoi(string(stone))
		if err != nil {
			panic(err)
		}
		out = append(out, v)
	}
	return out
}

func mutate(i int) (int, int) {
	if i == 0 {
		return 1, -1
	}
	if s := strconv.Itoa(i); len(s)%2 == 0 {
		h := len(s) / 2
		s1, s2 := s[:h], s[h:]
		i1, err := strconv.Atoi(s1)
		if err != nil {
			panic(err)
		}
		i2, err := strconv.Atoi(s2)
		if err != nil {
			panic(err)
		}
		return i1, i2
	}
	return i * 2024, -1
}

func p1(p []int) {
	puzzle := slices.Clone(p)
	for range 25 {
		for i := len(puzzle) - 1; i >= 0; i-- {
			s1, s2 := mutate(puzzle[i])
			puzzle[i] = s1
			if s2 != -1 {
				puzzle = slices.Insert(puzzle, i+1, s2)
			}
		}
	}
	fmt.Println(len(puzzle))
}
