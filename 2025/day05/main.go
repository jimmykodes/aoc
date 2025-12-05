package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	puzzle := parse("assets/input.txt")
	p1 := 0
	for _, item := range puzzle.items {
		for _, r := range puzzle.inventory {
			if r.contains(item) {
				p1++
				break
			}
		}
	}
	// fmt.Printf("p1 %d\n", p1)

	var p2 int64
	for _, r := range puzzle.inventory {
		p2 += r.size()
	}
	fmt.Printf("p2 %20d\n", (p2))
}

type Range struct {
	start int64
	end   int64
}

func (r Range) contains(v int64) bool {
	return r.start <= v && v <= r.end
}

func (r Range) size() int64 {
	return (r.end - r.start) + 1
}

func (r Range) overlaps(other Range) bool {
	switch {
	case other.start <= r.start && r.start <= other.end:
		// range starts in between other's start and end
		return true
	case other.start <= r.end && r.end <= other.end:
		// range ends in between other's start and end
		return true
	case r.start == other.end+1:
		// range starts immediately after other's end
		return true
	case r.end == other.start-1:
		// range ends immediately before other's start
		return true
	default:
		return false
	}
}

func (r Range) grow(other Range) Range {
	return Range{
		start: min(r.start, other.start),
		end:   max(r.end, other.end),
	}
}

type Puzzle struct {
	inventory []Range
	items     []int64
}

func parse(fname string) Puzzle {
	data, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	inventoryStr, itemsStr, _ := strings.Cut(string(data), "\n\n")

	inventoryLines := strings.Split(strings.TrimSpace(inventoryStr), "\n")
	inventory := make([]Range, len(inventoryLines))

	for i, line := range inventoryLines {
		startStr, endStr, _ := strings.Cut(line, "-")
		start, _ := strconv.ParseInt(startStr, 10, 64)
		end, _ := strconv.ParseInt(endStr, 10, 64)
		inventory[i] = Range{start: start, end: end}
	}

	slices.SortFunc(inventory, func(a, b Range) int {
		if a.start < b.start {
			return -1
		}
		return 1
	})
	inventory = completelyConslidate(inventory)

	for _, rng := range inventory {
		fmt.Printf("%s %s\n", formatWithSep(rng.start), formatWithSep(rng.end))
	}

	itemLines := strings.Split(strings.TrimSpace(itemsStr), "\n")
	items := make([]int64, len(itemLines))
	for i, line := range itemLines {
		item, _ := strconv.ParseInt(line, 10, 64)
		items[i] = item
	}

	return Puzzle{inventory: inventory, items: items}
}

func completelyConslidate(ranges []Range) []Range {
	out := ranges

	for {
		inner := conslidateRanges(out)
		if len(inner) == len(out) {
			return inner
		}
		out = inner
	}
}

func conslidateRanges(ranges []Range) []Range {
	var out []Range

	for _, rng := range ranges {
		var consolidated bool
		for i, other := range out {
			if rng.overlaps(other) {
				out[i] = rng.grow(other)
				consolidated = true
				break
			}
		}

		if !consolidated {
			out = append(out, rng)
		}
	}

	return out
}

func formatWithSep(n int64) string {
	str := strconv.FormatInt(n, 10)
	if len(str) <= 3 {
		return str
	}

	var result strings.Builder
	for i, digit := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result.WriteString("")
		}
		result.WriteRune(digit)
	}
	return result.String()
}
