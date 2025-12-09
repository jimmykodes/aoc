package main

import (
	"bytes"
	"errors"
	"fmt"
	"maps"
	"math/rand"
	"os"
	"slices"
	"sort"
	"strconv"
)

func main() {
	P2("assets/input.txt")
	// P1("assets/input.txt", 1_000)
}

func P2(fname string) {
	boxes := parse(fname)

	edges := make([]Edge, 0, len(boxes)*len(boxes))
	for i, box := range boxes {
		for _, otherBox := range boxes[i+1:] {
			if box == otherBox {
				continue
			}
			edges = append(edges, Edge{box, otherBox})
		}
	}

	slices.SortStableFunc(edges, func(a, b Edge) int {
		aDist := a.MagDistance()
		bDist := b.MagDistance()
		switch {
		case aDist < bDist:
			return -1
		case aDist > bDist:
			return 1
		default:
			return 0
		}
	})

	ptr := 0
	for len(circuitMap) > 1 {
		edge := edges[ptr]
		ptr++
		if edge[0].Circuit == edge[1].Circuit {
			// boxes are already on the same circuit
			continue
		}
		newCircuit := CombineCircuits(edge[0].Circuit, edge[1].Circuit)
		for _, box := range newCircuit.Boxes {
			box.Circuit = newCircuit
		}
	}
	last := edges[ptr-1]
	fmt.Println(last[0].X * last[1].X)
}

func P1(fname string, iterations int) {
	boxes := parse(fname)

	edges := make([]Edge, 0, len(boxes)*len(boxes))
	for i, box := range boxes {
		for _, otherBox := range boxes[i+1:] {
			if box == otherBox {
				continue
			}
			edges = append(edges, Edge{box, otherBox})
		}
	}
	fmt.Println(len(edges))

	slices.SortStableFunc(edges, func(a, b Edge) int {
		aDist := a.MagDistance()
		bDist := b.MagDistance()
		switch {
		case aDist < bDist:
			return -1
		case aDist > bDist:
			return 1
		default:
			return 0
		}
	})

	ptr := 0
	for ; ptr < iterations; ptr++ {
		edge := edges[ptr]
		if edge[0].Circuit == edge[1].Circuit {
			// boxes are already on the same circuit
			continue
		}
		newCircuit := CombineCircuits(edge[0].Circuit, edge[1].Circuit)
		for _, box := range newCircuit.Boxes {
			box.Circuit = newCircuit
		}
	}

	circuitSet := make(map[int64]*Circuit)
	for _, box := range boxes {
		circuitSet[box.Circuit.ID] = box.Circuit
	}

	circuits := slices.Collect(maps.Values(circuitSet))
	fmt.Println(len(circuits), "circuits")
	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i].Boxes) > len(circuits[j].Boxes)
	})

	product := 1
	for i := range 3 {
		product *= len(circuits[i].Boxes)
	}
	fmt.Println(product)

	for len(circuitMap) > 1 {
		edge := edges[ptr]
		ptr++
		if edge[0].Circuit == edge[1].Circuit {
			// boxes are already on the same circuit
			continue
		}
		newCircuit := CombineCircuits(edge[0].Circuit, edge[1].Circuit)
		for _, box := range newCircuit.Boxes {
			box.Circuit = newCircuit
		}
	}
	last := edges[ptr-1]
	fmt.Println(last[0].X * last[1].X)
}

type Circuit struct {
	ID    int64
	Boxes []*Box
}

var circuitMap = make(map[int64]*Circuit)

func NewCircuit(boxes []*Box) *Circuit {
	c := Circuit{
		ID:    rand.Int63(),
		Boxes: boxes,
	}
	circuitMap[c.ID] = &c
	return &c
}

func CombineCircuits(a, b *Circuit) *Circuit {
	delete(circuitMap, a.ID)
	delete(circuitMap, b.ID)
	return NewCircuit(append(a.Boxes, b.Boxes...))
}

type Edge [2]*Box

func (e Edge) MagDistance() int {
	dX := e[0].X - e[1].X
	dY := e[0].Y - e[1].Y
	dZ := e[0].Z - e[1].Z
	return (dX * dX) + (dY * dY) + (dZ * dZ)
}

func (e Edge) String() string {
	return fmt.Sprintf("%s -> %s", e[0], e[1])
}

// Eq returns true when two edges are equal
//
// Edges are considered equal if they contain the same junction
// boxes at either end. Directionality does not matter
func (e Edge) Eq(other Edge) bool {
	return (e[0] == other[0] && e[1] == other[1]) ||
		(e[1] == other[0] && e[0] == other[1])
}

type Box struct {
	X, Y, Z int
	Circuit *Circuit
}

func (b Box) String() string {
	return fmt.Sprintf("( %d, %d, %d )", b.X, b.Y, b.Z)
}

func (b *Box) UnmarshalText(data []byte) error {
	x, right, _ := bytes.Cut(data, []byte{','})
	y, z, _ := bytes.Cut(right, []byte{','})

	X, xErr := strconv.Atoi(string(x))
	Y, yErr := strconv.Atoi(string(y))
	Z, zErr := strconv.Atoi(string(z))
	b.X, b.Y, b.Z = X, Y, Z

	return errors.Join(xErr, yErr, zErr)
}

func parse(fname string) []*Box {
	data, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(bytes.TrimSpace(data), []byte{'\n'})
	out := make([]*Box, len(lines))
	for i, line := range lines {
		var b Box
		if err := b.UnmarshalText(line); err != nil {
			panic(err)
		}
		b.Circuit = NewCircuit([]*Box{&b})
		out[i] = &b
	}
	return out
}
