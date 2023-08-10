package main

import (
	_ "embed"
	"fmt"
	"sort"
)

var (
	//go:embed test.txt
	test string
	//go:embed input.txt
	input string
)

func main() {
	fmt.Println(p1())
	fmt.Println(p2())
}

func p1() int {
	monkeys := getMonkeys()
	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			monkey.Inspect(true)
		}
	}
	monkeyBusiness := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		monkeyBusiness[i] = monkey.inpectionCount
	}
	sort.Ints(monkeyBusiness)
	l := len(monkeys)
	return monkeyBusiness[l-1] * monkeyBusiness[l-2]
}

func p2() int {
	monkeys := getMonkeys()
	for i := 0; i < 10_000; i++ {
		for _, monkey := range monkeys {
			monkey.Inspect(false)
		}
	}
	monkeyBusiness := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		monkeyBusiness[i] = monkey.inpectionCount
	}
	sort.Ints(monkeyBusiness)
	l := len(monkeys)
	return monkeyBusiness[l-1] * monkeyBusiness[l-2]
}

func testFactory(monkeys []*Monkey) func(d, idx1, idx2 int) func(i int) *Monkey {
	return func(d, idx1, idx2 int) func(i int) *Monkey {
		return func(i int) *Monkey {
			if i%d == 0 {
				return monkeys[idx1]
			}
			return monkeys[idx2]
		}
	}
}

func testMonkeys() []*Monkey {
	mod := 23 * 19 * 13 * 17
	monkeys := []*Monkey{
		NewMonkey(0, func(i int) int { return i * 19 }, mod, 79, 98),
		NewMonkey(1, func(i int) int { return i + 6 }, mod, 54, 65, 75, 74),
		NewMonkey(2, func(i int) int { return i * i }, mod, 79, 60, 97),
		NewMonkey(3, func(i int) int { return i + 3 }, mod, 74),
	}
	test := testFactory(monkeys)
	monkeys[0].test = test(23, 2, 3)
	monkeys[1].test = test(19, 2, 0)
	monkeys[2].test = test(13, 1, 3)
	monkeys[3].test = test(17, 0, 1)
	return monkeys
}

func getMonkeys() []*Monkey {
	mod := 7 * 3 * 2 * 11 * 17 * 5 * 13 * 19
	m := []*Monkey{
		NewMonkey(0, func(i int) int { return i * 13 }, mod, 91, 58, 52, 69, 95, 54),
		NewMonkey(1, func(i int) int { return i * i }, mod, 80, 80, 97, 84),
		NewMonkey(2, func(i int) int { return i + 7 }, mod, 86, 92, 71),
		NewMonkey(3, func(i int) int { return i + 4 }, mod, 96, 90, 99, 76, 79, 85, 98, 61),
		NewMonkey(4, func(i int) int { return i * 19 }, mod, 60, 83, 68, 64, 73),
		NewMonkey(5, func(i int) int { return i + 3 }, mod, 96, 52, 52, 94, 76, 51, 57),
		NewMonkey(6, func(i int) int { return i + 5 }, mod, 75),
		NewMonkey(7, func(i int) int { return i + 1 }, mod, 83, 75),
	}

	test := testFactory(m)
	m[0].test = test(7, 1, 5)
	m[1].test = test(3, 3, 5)
	m[2].test = test(2, 0, 4)
	m[3].test = test(11, 7, 6)
	m[4].test = test(17, 1, 0)
	m[5].test = test(5, 7, 3)
	m[6].test = test(13, 4, 2)
	m[7].test = test(19, 2, 6)

	return m
}

func NewMonkey(id int, op func(int) int, divisor int, items ...int) *Monkey {
	return &Monkey{
		id:      id,
		items:   items,
		op:      op,
		divisor: divisor,
	}
}

type Monkey struct {
	id      int
	items   []int
	op      func(int) int
	test    func(int) *Monkey
	divisor int

	inpectionCount int
}

func (m *Monkey) Op(item int) int {
	return m.op(item) % m.divisor
}

func (m *Monkey) Add(item int) {
	m.items = append(m.items, item)
}

func (m *Monkey) Inspect(div bool) {
	for _, item := range m.items {
		item = m.Op(item)
		if div {
			item = item / 3
		}
		m.test(item).Add(item)
	}
	m.inpectionCount += len(m.items)
	m.items = nil
}
