package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	towels, patterns := puzzle("assets/input.txt")
	p1(towels, patterns)
	p2(towels, patterns)
}

func p1(towels Towels, patterns []string) {
	var total int
	for _, pattern := range patterns {
		if checkPattern(pattern, towels, 0) {
			total++
		}
	}
	fmt.Println(total)
}

func p2(towels Towels, patterns []string) {
	var total int
	for _, pattern := range patterns {
		total += countPattern(pattern, towels)
	}
	fmt.Println(total)
}

var counts = make(map[string]int)

func countPattern(pattern string, towels Towels) int {
	if res, ok := counts[pattern]; ok {
		return res
	}
	counts[pattern] = _countPattern(pattern, towels)
	return counts[pattern]
}

func _countPattern(pattern string, towels Towels) int {
	if pattern == "" {
		return 1
	}
	options := tokens(pattern, towels)
	if len(options) == 0 {
		return 0
	}

	var count int
	for _, option := range options {
		count += countPattern(pattern[len(option):], towels)
	}

	return count
}

var checked = make(map[string]bool)

func checkPattern(pattern string, towels Towels, inset int) bool {
	if res, ok := checked[pattern]; ok {
		return res
	}
	checked[pattern] = _checkPattern(pattern, towels, inset)
	return checked[pattern]
}

func _checkPattern(pattern string, towels Towels, inset int) bool {
	// prefix := strings.Repeat("|", inset)
	// fmt.Println(prefix, pattern)
	if pattern == "" {
		return true
	}

	options := tokens(pattern, towels)
	// fmt.Println(prefix, options)
	if len(options) == 0 {
		return false
	}

	for _, option := range options {
		if checkPattern(pattern[len(option):], towels, inset+1) {
			return true
		}
	}

	return false
}

func tokens(pattern string, towels Towels) []string {
	var out []string
	rp := len(pattern)
	for {
		check := pattern[:rp]
		if _, ok := towels[check]; ok {
			out = append(out, check)
		}
		rp--
		if rp < 0 {
			return out
		}
	}
}

func verboseCheck(pattern string, towels Towels) {
	lp := 0
	for lp < len(pattern) {
		rp := len(pattern)
		for {
			check := pattern[lp:rp]
			if _, ok := towels[string(check)]; ok {
				fmt.Println("matched", check)
				lp += len(check)
				break
			} else {
				fmt.Println(check, "failed")
				rp--
				if rp == lp {
					fmt.Println("final fail")
					return
				}
			}
		}
	}
}

type (
	Towels map[string]struct{}
)

func puzzle(filename string) (Towels, []string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	towelData, patternData, _ := bytes.Cut(data, []byte{'\n', '\n'})
	towels := make(Towels)
	for _, towel := range bytes.Split(towelData, []byte{',', ' '}) {
		towels[string(towel)] = struct{}{}
	}
	var patterns []string
	for _, pattern := range bytes.Fields(patternData) {
		patterns = append(patterns, string(pattern))
	}
	return towels, patterns
}
