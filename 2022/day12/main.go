package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strings"
)

var (
	//go:embed input.txt
	input string
	//go:embed test.txt
	test string
)

func NewCoord(x, y int) *Coord {
	return &Coord{
		x:    x,
		y:    y,
		dist: math.MaxInt64,
	}
}

type Coord struct {
	x, y   int
	height rune

	from     *Coord
	finished bool
	dist     int64

	left   *Coord
	right  *Coord
	top    *Coord
	bottom *Coord
}

func (c *Coord) SetLeft(other *Coord) {
	c.left = other
	other.right = c
}
func (c *Coord) SetTop(other *Coord) {
	c.top = other
	other.bottom = c
}

func (c *Coord) String() string {
	return fmt.Sprintf("<Coord: (%d, %d) - %3d - %2d>", c.x, c.y, c.height, c.dist)
}

type Board [][]*Coord

func getData() (*Coord, *Coord, Board) {
	var (
		start, end *Coord
	)
	lines := strings.Split(input, "\n")
	out := make([][]*Coord, len(lines))
	for y, line := range lines {
		row := make([]*Coord, len(line))
		for x, c := range line {
			coord := NewCoord(x, y)
			switch c {
			case 'S':
				coord.height = 'a'
				start = coord
			case 'E':
				coord.height = 'z'
				end = coord
			default:
				coord.height = c
			}
			row[x] = coord
			if x > 0 {
				coord.SetLeft(row[x-1])
			}
			if y > 0 {
				coord.SetTop(out[y-1][x])
			}
		}
		out[y] = row
	}
	return start, end, out
}

func canStep(from, to *Coord) bool {
	return to.height-from.height <= 1
}

func NewQueue() *Queue {
	return &Queue{seen: make(map[string]struct{})}
}

type Queue struct {
	coords []*Coord
	seen   map[string]struct{}
}

func (q *Queue) Len() int           { return len(q.coords) }
func (q *Queue) Less(i, j int) bool { return q.coords[i].dist < q.coords[j].dist }
func (q *Queue) Swap(i, j int)      { q.coords[i], q.coords[j] = q.coords[j], q.coords[i] }

func (q *Queue) Add(c *Coord) {
	key := fmt.Sprintf("%d-%d", c.x, c.y)
	if _, ok := q.seen[key]; ok {
		// already added, do nothing
		return
	}
	q.seen[key] = struct{}{}
	q.coords = append(q.coords, c)
}

func (q *Queue) Pop() *Coord {
	c := q.coords[0]
	q.coords = q.coords[1:]
	return c
}

func main() {
	fmt.Println(p1())
	fmt.Println(p2())
}

func p1() int64 {
	start, end, _ := getData()
	start.dist = 0
	queue := NewQueue()
	queue.Add(start)
	for queue.Len() > 0 {
		node := queue.Pop()
		node.finished = true
		dist := node.dist + 1
		for _, n := range []*Coord{node.left, node.right, node.top, node.bottom} {
			if n != nil && canStep(node, n) && !n.finished {
				if dist < n.dist {
					// faster to get to l through current node than previous `from` node
					n.from = node
					n.dist = dist
				}
				queue.Add(n)
			}
		}
		sort.Sort(queue)
	}

	if !end.finished {
		panic("we didn't finish")
	}
	return end.dist
}

func p2() int64 {
	_, start, _ := getData()
	start.dist = 0
	queue := NewQueue()
	queue.Add(start)
	for queue.Len() > 0 {
		node := queue.Pop()
		if node.height == 'a' {
			return node.dist
		}
		node.finished = true
		dist := node.dist + 1
		for _, n := range []*Coord{node.left, node.right, node.top, node.bottom} {
			if n != nil && canStep(n, node) && !n.finished {
				if dist < n.dist {
					// faster to get to l through current node than previous `from` node
					n.from = node
					n.dist = dist
				}
				queue.Add(n)
			}
		}
		sort.Sort(queue)
	}
	return math.MaxInt64
}
