package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

var fname = "assets/input.txt"

func main() {
	start := time.Now()
	P2()
	fmt.Printf("time.Since(start): %v\n", time.Since(start))
}

func P2() {
	machines := parse(fname)
	var total int
	for i, m := range machines {
		total += m.FindJoltage()
		fmt.Printf("i: %v\n", i)
	}
	fmt.Printf("total: %v\n", total)
}

func P1() {
	machines := parse(fname)
	var total int
	for _, m := range machines {
		total += m.TargetBFS()
	}
	fmt.Printf("total: %v\n", total)
}

type Machine struct {
	Target  int
	Buttons []int
	Joltage []int
}

func (m Machine) FindJoltage() int {
	type queueItem struct {
		count    int
		joltages []int
	}
	maxJoltage := 0
	for _, joltage := range m.Joltage {
		maxJoltage = max(maxJoltage, joltage)
	}
	queue := []*queueItem{
		{count: 0, joltages: make([]int, len(m.Joltage))},
	}
	var item *queueItem
	for len(queue) > 0 {
		item, queue = queue[0], queue[1:]
		for _, button := range m.Buttons {
			newJoltages := joltagePress(button, item.joltages)
			if item.count >= maxJoltage {
				// only compare if we could possibly have hit it
				switch compareJoltages(newJoltages, m.Joltage) {
				case 0:
					return item.count + 1
				case -1:
					i := queueItem{count: item.count + 1, joltages: newJoltages}
					queue = append(queue, &i)
				}
			} else {
				i := queueItem{count: item.count + 1, joltages: newJoltages}
				queue = append(queue, &i)
			}
		}
	}
	panic("nothing found")
}

// return 0 for equal, -1 for all less, 1 for any greater
func compareJoltages(have, want []int) int {
	var notEqual int
	for i, v := range have {
		if want[i] < v {
			return 1
		}
		if want[i] > v {
			notEqual = -1
		}
	}
	return notEqual
}

func joltagePress(button int, joltages []int) []int {
	out := slices.Clone(joltages)
	for i := range out {
		if button&(1<<i) > 0 {
			out[i]++
		}
	}
	return out
}

func (m Machine) TargetBFS() int {
	queue := [][2]int{{0, 0}}
	for len(queue) > 0 {
		var v [2]int
		v, queue = queue[0], queue[1:]
		val, count := v[0], v[1]
		for _, button := range m.Buttons {
			n := [2]int{val ^ button, count + 1}
			if n[0] == m.Target {
				return n[1]
			}
			queue = append(queue, n)
		}
	}
	panic("never found value")
}

func parse(fname string) []*Machine {
	data, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(bytes.TrimSpace(data), []byte{'\n'})
	out := make([]*Machine, len(lines))
	for i, line := range lines {
		parts := strings.Split(string(line), " ")
		target := parts[0]
		buttons := parts[1 : len(parts)-1]
		joltage := parts[len(parts)-1]
		out[i] = &Machine{
			Target:  parseTarget(target),
			Buttons: parseButtons(buttons),
			Joltage: parseJoltage(joltage),
		}
	}
	return out
}

func parseJoltage(s string) []int {
	s = s[1 : len(s)-1] // trim open and close braces
	nums := strings.Split(s, ",")

	out := make([]int, len(nums))
	for i, numStr := range nums {
		val, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		out[i] = val
	}

	return out
}

func parseTarget(s string) int {
	s = s[1 : len(s)-1] // trim open and close brackets
	var out int

	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			out |= 1 << i
		}
	}
	return out
}

func parseButtons(buttons []string) []int {
	out := make([]int, len(buttons))
	for i, button := range buttons {
		out[i] = parseButton(button)
	}
	return out
}

func parseButton(s string) int {
	s = s[1 : len(s)-1] // trim open and close parens
	nums := strings.Split(s, ",")

	var out int
	for _, numStr := range nums {
		idx, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		out |= 1 << idx
	}

	return out
}
