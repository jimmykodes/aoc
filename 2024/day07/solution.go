package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	puzzles := parse("assets/input.txt")
	p2(puzzles)
}

type Puzzle struct {
	total int
	terms []int
}

func parse(filename string) []*Puzzle {
	data, err := os.ReadFile(filename)
	data = bytes.TrimSpace(data)
	if err != nil {
		panic(err)
	}
	var out []*Puzzle
	for _, line := range bytes.Split(data, []byte{'\n'}) {
		totalBytes, termBytes, _ := bytes.Cut(line, []byte(": "))

		p := Puzzle{}
		p.total, err = strconv.Atoi(string(totalBytes))
		if err != nil {
			panic(err)
		}

		for _, termStr := range strings.Split(string(termBytes), " ") {
			term, err := strconv.Atoi(termStr)
			if err != nil {
				panic(err)
			}
			p.terms = append(p.terms, term)
		}
		out = append(out, &p)
	}
	return out
}

type Node struct {
	value  int
	left   *Node
	middle *Node
	right  *Node
}

func p2(puzzles []*Puzzle) {
	var total int
	for _, puzzle := range puzzles {

		var start *Node

		// build tree
		var leaves []*Node
		for i, term := range puzzle.terms {
			if i == 0 {
				n := &Node{value: term}
				start = n
				leaves = append(leaves, n)
				continue
			}
			var _leaves []*Node
			for _, leaf := range leaves {
				leaf.left = &Node{value: term}
				leaf.right = &Node{value: term}
				leaf.middle = &Node{value: term}
				_leaves = append(_leaves, leaf.left, leaf.right, leaf.middle)
			}
			leaves = _leaves
		}

		if check(start, puzzle.total) {
			total += puzzle.total
		}
	}
	fmt.Println(total)
}

func p1(puzzles []*Puzzle) {
	var total int
	for _, puzzle := range puzzles {

		var start *Node

		// build tree
		var leaves []*Node
		for i, term := range puzzle.terms {
			if i == 0 {
				n := &Node{value: term}
				start = n
				leaves = append(leaves, n)
				continue
			}
			var _leaves []*Node
			for _, leaf := range leaves {
				leaf.left = &Node{value: term}
				leaf.right = &Node{value: term}
				_leaves = append(_leaves, leaf.left, leaf.right)
			}
			leaves = _leaves
		}

		if check(start, puzzle.total) {
			total += puzzle.total
		}
	}
	fmt.Println(total)
}

func check(start *Node, target int) bool {
	stack := []*Node{start}
	for len(stack) > 0 {
		var n *Node
		n, stack = stack[0], stack[1:]
		if n.value > target {
			continue
		}
		if n.left == nil {
			// on a leaf node - check for target value
			if n.value == target {
				return true
			}
			continue
		}
		n.left.value = n.value + n.left.value
		n.right.value = n.value * n.right.value
		stack = append(stack, n.left, n.right)
		if n.middle != nil {
			// in part two
			var err error
			n.middle.value, err = strconv.Atoi(strconv.Itoa(n.value) + strconv.Itoa(n.middle.value))
			if err != nil {
				panic(err)
			}
			stack = append(stack, n.middle)
		}

	}
	return false
}
