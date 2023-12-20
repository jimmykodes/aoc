package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("assets/input.txt")
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	fmt.Println(p2(data))
}

func p1(data []byte) int {
	current := 0
	for _, chunk := range bytes.Split(data, []byte(",")) {
		current += hash(chunk)
	}
	return current
}

type Operation int

const (
	RemoveOperation Operation = iota
	SetOperation
)

type Lens struct {
	Label     []byte
	Length    int
	Operation Operation
}

func (l *Lens) UnmarshalText(data []byte) error {
	label, length, found := bytes.Cut(data, []byte{'='})
	if found {
		l.Label = label
		l.Length, _ = strconv.Atoi(string(length))
		l.Operation = SetOperation
		return nil
	}

	l.Label = bytes.TrimSuffix(data, []byte{'-'})
	return nil
}

type Box struct {
	Lenses []Lens
}

func (b *Box) Add(l Lens) {
	for i, lens := range b.Lenses {
		if bytes.Equal(lens.Label, l.Label) {
			b.Lenses[i] = l
			return
		}
	}
	b.Lenses = append(b.Lenses, l)
}

func (b *Box) Remove(l Lens) {
	for i, lens := range b.Lenses {
		if bytes.Equal(lens.Label, l.Label) {
			b.Lenses = append(b.Lenses[:i], b.Lenses[i+1:]...)
			return
		}
	}
}

func p2(data []byte) int {
	var boxes [256]Box
	for _, chunk := range bytes.Split(data, []byte(",")) {
		var l Lens
		_ = l.UnmarshalText(chunk)
		boxNum := hash(l.Label)
		switch l.Operation {
		case SetOperation:
			boxes[boxNum].Add(l)
		case RemoveOperation:
			boxes[boxNum].Remove(l)
		}
	}
	total := 0
	for i, box := range boxes {
		for j, lens := range box.Lenses {
			total += (i + 1) * (j + 1) * lens.Length
		}
	}
	return total
}

func hash(input []byte) int {
	var current int
	for _, b := range input {
		current += int(b)
		current *= 17
		current %= 256
	}
	return current
}
