package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	dirs, nodeMap := load("assets/input.txt")
	// dirs, nodeMap := load("assets/test2.txt")
	// fmt.Println(p1(dirs, nodeMap["AAA"]))
	fmt.Println(p2(dirs, nodeMap))
}

func p1(dirs []Direction, node *Node) int {
	dirIdx := 0
	steps := 0
	for node.ID != "ZZZ" {
		steps++
		node = node.Next(dirs[dirIdx])
		dirIdx = (dirIdx + 1) % len(dirs)
	}
	return steps
}

func p2(dirs []Direction, nodeMap map[string]*Node) int {
	var nodes []*Node
	for id, node := range nodeMap {
		if id[2] == 'A' {
			nodes = append(nodes, node)
		}
	}

	periods := make([]int, len(nodes))
	for i, node := range nodes {
		periods[i] = findPeriod(dirs, node)
	}
	fmt.Println("periods:")
	fmt.Println(periods)
	fmt.Println("is it bad i used wolfram alpha to find the LCM of the periods above?")
	fmt.Println("i should probably figure out how to do it in code... but i'm lazy")
	return -1
}

func findPeriod(dirs []Direction, node *Node) int {
	var steps, dir int
	for node.ID[2] != 'Z' {
		node = node.Next(dirs[dir])
		steps++
		dir++
		dir %= len(dirs)
	}
	return steps
}

type Direction int

const (
	Left Direction = iota + 1
	Right
)

func load(fn string) ([]Direction, map[string]*Node) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	dirStr := scanner.Text()
	scanner.Scan()
	_ = scanner.Bytes()

	dirs := make([]Direction, len(dirStr))
	for i, r := range dirStr {
		switch r {
		case 'L':
			dirs[i] = Left
		case 'R':
			dirs[i] = Right
		default:
			panic("invalid run")
		}
	}

	nodeMap := make(map[string]*Node)
	for scanner.Scan() {
		line := scanner.Text()
		nodeID, sides, _ := strings.Cut(line, " = ")
		node := nodeMap[nodeID]
		if node == nil {
			nodeMap[nodeID] = NewNode(nodeID)
			node = nodeMap[nodeID]
		}
		leftID, rightID, _ := strings.Cut(sides, ", ")
		leftID = strings.TrimLeft(leftID, "(")
		rightID = strings.TrimRight(rightID, ")")

		left := nodeMap[leftID]
		if left == nil {
			nodeMap[leftID] = NewNode(leftID)
			left = nodeMap[leftID]
		}

		right := nodeMap[rightID]
		if right == nil {
			nodeMap[rightID] = NewNode(rightID)
			right = nodeMap[rightID]
		}

		node.Left = left
		node.Right = right
	}
	return dirs, nodeMap
}

func NewNode(id string) *Node {
	return &Node{ID: id}
}

type Node struct {
	ID          string
	Left, Right *Node
}

func (n Node) Next(dir Direction) *Node {
	switch dir {
	case Left:
		return n.Left
	case Right:
		return n.Right
	}
	return nil
}
