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
	p2(p)
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

var seen = map[[2]int]int{}

func mutate(stone int) []int {
	if stone == 0 {
		return []int{1}
	}
	if s := strconv.Itoa(stone); len(s)%2 == 0 {
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
		return []int{i1, i2}
	}
	return []int{stone * 2024}
}

func count(stone, generations int) int {
	if generations == 0 {
		return 1
	}
	cacheKey := [2]int{stone, generations}
	if c, ok := seen[cacheKey]; ok {
		return c
	}
	var c int
	for _, s := range mutate(stone) {
		c += count(s, generations-1)
	}
	seen[cacheKey] = c
	return c
}

// _p1 is my initial, stupid, brute force attempt
func _p1(p []int) {
	puzzle := slices.Clone(p)
	for range 25 {
		for i := len(puzzle) - 1; i >= 0; i-- {
			vals := mutate(puzzle[i])
			puzzle[i] = vals[0]
			if len(vals) > 1 {
				puzzle = slices.Insert(puzzle, i+1, vals[1])
			}
		}
	}
	fmt.Println(len(puzzle))
}

func solve(p []int, gens int) int {
	puzzle := slices.Clone(p)
	var total int
	for _, stone := range puzzle {
		total += count(stone, gens)
	}
	return total
}

func p1(p []int) {
	fmt.Println(solve(p, 25))
}

func p2(p []int) {
	fmt.Println(solve(p, 75))
}
