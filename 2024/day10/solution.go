package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	th, _ := puzzle("assets/input.txt")
	p1(th)
	p2(th)
}

func p1(trailheads []*Node) {
	var total int
	for _, th := range trailheads {
		stack := []*Node{th}
		peaks := map[string]struct{}{}
		for len(stack) > 0 {
			var node *Node
			node, stack = stack[0], stack[1:]
			if node.Height == 9 {
				peaks[node.ID] = struct{}{}
				continue
			}
			if node.Up != nil && node.Up.Height-node.Height == 1 {
				stack = append(stack, node.Up)
			}
			if node.Down != nil && node.Down.Height-node.Height == 1 {
				stack = append(stack, node.Down)
			}
			if node.Left != nil && node.Left.Height-node.Height == 1 {
				stack = append(stack, node.Left)
			}
			if node.Right != nil && node.Right.Height-node.Height == 1 {
				stack = append(stack, node.Right)
			}
		}
		total += len(peaks)
	}
	fmt.Println(total)
}

func p2(trailheads []*Node) {
	var total int
	for _, th := range trailheads {
		stack := []*Node{th}
		peaks := map[string]int{}
		for len(stack) > 0 {
			var node *Node
			node, stack = stack[0], stack[1:]
			if node.Height == 9 {
				peaks[node.ID]++
				continue
			}
			if node.Up != nil && node.Up.Height-node.Height == 1 {
				stack = append(stack, node.Up)
			}
			if node.Down != nil && node.Down.Height-node.Height == 1 {
				stack = append(stack, node.Down)
			}
			if node.Left != nil && node.Left.Height-node.Height == 1 {
				stack = append(stack, node.Left)
			}
			if node.Right != nil && node.Right.Height-node.Height == 1 {
				stack = append(stack, node.Right)
			}
		}
		for _, rating := range peaks {
			total += rating
		}
	}
	fmt.Println(total)
}

type Node struct {
	ID                    string // "X:Y"
	Height                int
	Up, Down, Left, Right *Node
}

func puzzle(filename string) ([]*Node, [][]*Node) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	charToHeight := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	data = bytes.TrimSpace(data)
	lines := bytes.Fields(data)
	var trailheads []*Node
	board := make([][]*Node, len(lines))
	for rowNum, line := range lines {
		row := make([]*Node, len(line))
		for colNum, char := range line {
			n := &Node{
				ID:     fmt.Sprintf("%d:%d", colNum, rowNum),
				Height: charToHeight[char],
			}
			if colNum > 0 {
				n.Left = row[colNum-1]
				row[colNum-1].Right = n
			}
			if rowNum > 0 {
				n.Up = board[rowNum-1][colNum]
				board[rowNum-1][colNum].Down = n
			}
			row[colNum] = n
			if n.Height == 0 {
				trailheads = append(trailheads, n)
			}
		}
		board[rowNum] = row
	}
	return trailheads, board
}
