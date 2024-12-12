package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	p := puzzle("assets/input.txt")
	// p1(p)
	p2(p)
}

func p1(puzzle [][]*Node) {
	var total int
	for _, region := range puzzle {
		perimeter := 0
		for _, node := range region {
			for _, p := range node.Perimeter() {
				if p {
					perimeter++
				}
			}
		}
		total += perimeter * len(region)
	}
	fmt.Println(total)
}

func _p2(puzzle [][]*Node) {
	var total int
	for _, region := range puzzle {
		switch len(region) {
		case 1:
			total += 4
		case 2:
			total += 8
		default:

			sides := 0
			node := region[0]
			visited := 0
			var d dir
			for {
				// due to the way we gather nodes in a region, we know that we will always start on a top left corner
				//
				// if perimeter[up] and not perimeter[right], keep going right until there is a right perimeter,
				// then turn and go down until we see while down and not right,
				// then turn again and so on, until we reach the start node.
				// This won't find interiors though... how do we do that?
				// reset nodes ?
				neighbor := node.Neighbors[d.turn()]
				for {
					fmt.Println(node, neighbor)
					if neighbor == nil {
						break
					}
					perimeters := neighbor.Perimeter()
					if neighbor.Label != node.Label || !perimeters[d] {
						break
					}
					node, neighbor = neighbor, neighbor.Neighbors[d.turn()]
					visited++
				}
				sides++
				if node == region[0] && visited > 1 {
					if p := node.Perimeter(); p[up] && p[down] && p[left] {
						sides++
					}
					break
				}
				if node.Perimeter()[d.turn()] {
					d = d.turn()
				} else {
					node = neighbor
					d = d.turn().turn().turn()
				}

			}

			fmt.Println("sides", sides, "area", len(region))
			total += sides * len(region)
		}
	}
	fmt.Println(total)
}

func p2(puzzle [][]*Node) {
	var total int
	for _, region := range puzzle {
		sides := 0
		for _, node := range region {
			sides += node.Corners()
		}
		total += sides * len(region)
	}
	fmt.Println(total)
}

type Node struct {
	X, Y      int
	Visited   bool
	Label     rune
	Neighbors [4]*Node
}

type dir int

func (d dir) turn() dir {
	return (d + 1) % mod
}

const (
	up dir = iota
	right
	down
	left
	mod
)

func (n Node) Perimeter() [4]bool {
	var out [4]bool
	for i, neighbor := range n.Neighbors {
		if neighbor == nil || neighbor.Label != n.Label {
			out[i] = true
		}
	}
	return out
}

func (n Node) Corners() int {
	p := n.Perimeter()
	out := 0
	// up and left
	if p[up] && p[left] {
		out++
	}
	if !p[up] && !p[left] {
		if n.Neighbors[up].Perimeter()[left] && n.Neighbors[left].Perimeter()[up] {
			out++
		}
	}
	// down and left
	if p[down] && p[left] {
		out++
	}
	if !p[down] && !p[left] {
		if n.Neighbors[down].Perimeter()[left] && n.Neighbors[left].Perimeter()[down] {
			out++
		}
	}
	// up and right
	if p[up] && p[right] {
		out++
	}
	if !p[up] && !p[right] {
		if n.Neighbors[up].Perimeter()[right] && n.Neighbors[right].Perimeter()[up] {
			out++
		}
	}

	// down and right
	if p[down] && p[right] {
		out++
	}
	if !p[down] && !p[right] {
		if n.Neighbors[down].Perimeter()[right] && n.Neighbors[right].Perimeter()[down] {
			out++
		}
	}

	return out
}

func (n Node) String() string {
	return fmt.Sprintf("<Node(%d, %d): Label=%q Visited=%v>", n.X, n.Y, n.Label, n.Visited)
}

func puzzle(filename string) [][]*Node {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := bytes.Fields(bytes.TrimSpace(data))
	board := make([][]*Node, len(lines))
	for rowNum, line := range lines {
		board[rowNum] = make([]*Node, len(line))
		for colNum, char := range line {
			n := &Node{X: colNum, Y: rowNum, Label: rune(char)}
			board[rowNum][colNum] = n
			if rowNum > 0 {
				n.Neighbors[up] = board[rowNum-1][colNum]
				board[rowNum-1][colNum].Neighbors[down] = n
			}
			if colNum > 0 {
				n.Neighbors[left] = board[rowNum][colNum-1]
				board[rowNum][colNum-1].Neighbors[right] = n
			}
		}
	}
	var regions [][]*Node
	for _, row := range board {
		for _, node := range row {
			if node.Visited {
				continue
			}
			node.Visited = true
			region := []*Node{node}
			stack := []*Node{node}
			for len(stack) > 0 {
				var n *Node
				n, stack = stack[0], stack[1:]
				for _, neighbor := range n.Neighbors {
					if neighbor == nil || neighbor.Visited {
						continue
					}
					if neighbor.Label == node.Label {
						region = append(region, neighbor)
						stack = append(stack, neighbor)
						neighbor.Visited = true
					}
				}
			}
			regions = append(regions, region)
		}
	}
	return regions
}
