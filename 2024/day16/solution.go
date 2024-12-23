package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	start, end, all := puzzle("assets/test.txt")
	stack := []Path{{Node: start, Dir: Left}}
	for len(stack) > 0 {
		path, stack := stack[0], stack[1:]
		for dir, neighbor := range path.Node.Neighbors {
			if neighbor == nil {
				continue
			}
		}
	}
}

type Path struct {
	Node *Node
	Dir  Dir
	Cost int
	From *Node
}

type Dir int

const (
	Up Dir = iota
	Right
	Down
	Left
	mod
)

func (d Dir) Flip() Dir {
	return (d + 2) % mod
}

type Node struct {
	X, Y      int
	Visited   bool
	Neighbors [4]*Node
}

func (n *Node) Link(d Dir, o *Node) bool {
	if o == nil {
		return false
	}
	n.Neighbors[d] = o
	o.Neighbors[d.Flip()] = n
	return true
}

func puzzle(filename string) (*Node, *Node, map[string]*Node) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	rows := bytes.Fields(data)
	above := make([]*Node, len(rows[0]))
	var left, start, end *Node
	out := make(map[string]*Node)

	for y, row := range rows {
		// new row, reset left
		left = nil
		for x, col := range row {
			switch col {
			case '.':
				n := &Node{X: x, Y: y}
				out[fmt.Sprintf("%s:%s", x, y)] = n
				switch {
				case row[x+1] == '#':
					// this is a junction or a dead end
					n.Link(Left, left)
				case left == nil, n.Link(Up, above[x]), rows[y+1][x] == '.':
					// left is nil
					// or we successfully linked above us
					// or we have an open space below us
					// we become the leftmost node
					left = n
				}
				above[x] = n
			case '#':
				above[x] = nil
				left = nil
			case 'S':
				start = &Node{X: x, Y: y}
				start.Link(Up, above[x])
				above[x] = start
				start.Link(Left, left)
				left = start
				out[fmt.Sprintf("%s:%s", x, y)] = start
			case 'E':
				end = &Node{X: x, Y: y}
				end.Link(Up, above[x])
				above[x] = end
				end.Link(Left, left)
				left = end
				out[fmt.Sprintf("%s:%s", x, y)] = end
			}
		}
	}

	return start, end, out
}
