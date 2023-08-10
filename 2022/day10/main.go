package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println(p1())
	p2()
}

func p1() int {
	c := &cpu{
		x:     1,
		stack: get_data(),
	}
	total := 0
	for c.next() {
		if c.cycle == 20 || (c.cycle-20)%40 == 0 {
			total += c.cycle * c.x
		}
		dx, done := c.op.tick()
		if done {
			c.op = nil
		}
		c.x += dx
	}
	return total
}

func p2() {
	c := &cpu{
		x:     1,
		stack: get_data(),
	}
	for c.next() {
		col := c.cycle % 40
		if math.Abs(float64(col-(c.x+1))) <= 1 {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
		if col == 0 {
			fmt.Println()
		}
		dx, done := c.op.tick()
		if done {
			c.op = nil
		}
		c.x += dx
	}
}

type op struct {
	count int
	value int
}

func (o *op) tick() (int, bool) {
	o.count--
	if o.count == 0 {
		return o.value, true
	}
	return 0, false
}

func noop() *op {
	return &op{count: 1}
}

func addx(v int) *op {
	return &op{count: 2, value: v}
}

type cpu struct {
	cycle int
	x     int
	stack []*op
	op    *op
}

func (c *cpu) next() bool {
	if len(c.stack) == 0 && c.op == nil {
		return false
	}
	c.cycle++
	if c.op == nil {
		c.op = c.stack[0]
		c.stack = c.stack[1:]
	}
	return true
}

func get_data() []*op {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(data, []byte("\n"))
	s := make([]*op, len(lines))
	for i, line := range lines {
		parts := bytes.SplitN(line, []byte(" "), 2)
		switch string(parts[0]) {
		case "noop":
			s[i] = noop()
		case "addx":
			dx, err := strconv.Atoi(string(parts[1]))
			if err != nil {
				panic(err)
			}
			s[i] = addx(dx)
		}
	}
	return s
}
