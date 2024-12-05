package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rules, updates := parsePuzzle("assets/input.txt")
	p1(rules, updates)
	p2(rules, updates)
}

func parsePuzzle(filename string) (map[string]*Rule, [][]string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	ruleData, updateData, _ := bytes.Cut(data, []byte("\n\n"))
	updateLines := bytes.Split(bytes.TrimSpace(updateData), []byte{'\n'})
	updates := make([][]string, len(updateLines))
	for i, updateLine := range updateLines {
		updates[i] = strings.Split(string(bytes.TrimSpace(updateLine)), ",")
	}

	ruleLines := bytes.Split(ruleData, []byte("\n"))
	rules := make(map[string]*Rule)
	for _, line := range ruleLines {
		before, after, _ := strings.Cut(string(line), "|")

		rb := rules[before]
		if rb == nil {
			rules[before] = new(Rule)
			rb = rules[before]
		}
		rb.appendAfter(after)

		ra := rules[after]
		if ra == nil {
			rules[after] = new(Rule)
			ra = rules[after]
		}
		ra.appendBefore(before)
	}
	return rules, updates
}

func p1(rules map[string]*Rule, updates [][]string) {
	var valid []int
	sf := sortFunc(rules)
	for j, update := range updates {
		if slices.IsSortedFunc(update, sf) {
			valid = append(valid, j)
		}
	}
	total := 0
	for _, idx := range valid {
		v := updates[idx]
		m, err := strconv.Atoi(v[len(v)/2])
		if err != nil {
			panic(err)
		}
		total += m
	}
	fmt.Println(total)
}

func p2(rules map[string]*Rule, updates [][]string) {
	var invalid []int
	sf := sortFunc(rules)

	for j, update := range updates {
		if !slices.IsSortedFunc(update, sf) {
			invalid = append(invalid, j)
		}
	}

	total := 0
	for _, idx := range invalid {
		v := updates[idx]
		slices.SortFunc(v, sf)
		m, err := strconv.Atoi(v[len(v)/2])
		if err != nil {
			panic(err)
		}
		total += m
	}
	fmt.Println(total)
}

func sortFunc(rules map[string]*Rule) func(a, b string) int {
	return func(a, b string) int {
		// -1 when a < b
		// 1 when a > b
		// 0 when a == b
		aRule := rules[a]
		if slices.Index(aRule.before, b) != -1 {
			// b is in a's 'before' so a > b
			return 1
		}
		// rules seem to be exhaustive, so it's reasonable to assume
		// that if it isn't in 'before' it is in 'after' so assume
		// a < b
		return -1
	}
}

type Rule struct {
	before []string
	after  []string
}

func (r *Rule) appendBefore(b string) {
	r.before = append(r.before, b)
}

func (r *Rule) appendAfter(b string) {
	r.after = append(r.after, b)
}

func (r Rule) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "<before: %+v - after: %+v>", r.before, r.after)
	return sb.String()
}
