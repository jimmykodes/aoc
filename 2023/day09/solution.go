package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	l := lines("assets/input.txt")
	fmt.Println(p1(l))
	fmt.Println(p2(l))
}

func p1(lines [][]int) int {
	total := 0
	for _, line := range lines {
		var stack [][]int
		stack = append(stack, line)
		for {
			l := stack[len(stack)-1]
			n := make([]int, len(l)-1)
			done := true
			for i := 1; i < len(l); i++ {
				delta := l[i] - l[i-1]
				if delta != 0 {
					// we encountered non-zero so we'll have to do this again
					done = false
				}
				n[i-1] = delta
			}
			if done {
				break
			}
			// no reason to add the 0 row, we know it'll always be there
			stack = append(stack, n)
		}
		for i := len(stack) - 1; i >= 0; i-- {
			var last_val int
			if i != len(stack)-1 {
				last_val = stack[i+1][len(stack[i+1])-1]
			}
			stack[i] = append(stack[i], stack[i][len(stack[i])-1]+last_val)
		}
		total += stack[0][len(stack[0])-1]
	}
	return total
}

func p2(lines [][]int) int {
	total := 0
	for _, line := range lines {
		var stack [][]int
		stack = append(stack, line)
		for {
			l := stack[len(stack)-1]
			n := make([]int, len(l)-1)
			done := true
			for i := 1; i < len(l); i++ {
				delta := l[i] - l[i-1]
				if delta != 0 {
					// we encountered non-zero so we'll have to do this again
					done = false
				}
				n[i-1] = delta
			}
			if done {
				break
			}
			// no reason to add the 0 row, we know it'll always be there
			stack = append(stack, n)
		}
		for i := len(stack) - 1; i >= 0; i-- {
			var last_val int
			if i != len(stack)-1 {
				last_val = stack[i+1][0]
			}
			stack[i] = append([]int{stack[i][0] - last_val}, stack[i]...)
		}
		total += stack[0][0]
	}
	return total
}

func lines(fn string) [][]int {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var l [][]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		out := make([]int, len(line))
		for i, num := range line {
			out[i], _ = strconv.Atoi(num)
		}
		l = append(l, out)
	}
	return l
}
