package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data, err := getData("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(p1(data))
	fmt.Println(p2(data))
}

func getData(filename string) ([]*move, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var moves []*move
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		move, err := newMove(scanner.Text())
		if err != nil {
			return nil, err
		}
		moves = append(moves, move)
	}
	return moves, scanner.Err()
}

func p1(moves []*move) int {
	head := newKnot(0)
	tail := newKnot(1)
	head.downstream = tail
	tail.upstream = head
	for _, move := range moves {
		for i := 0; i < move.distance; i++ {
			head.move(move.direction)
		}
	}
	total := 0
	for _, m := range tail.seen {
		total += len(m)
	}
	return total
}

func p2(moves []*move) int {
	head := newKnot(0)
	upstream := head
	for i := 1; i < 10; i++ {
		k := newKnot(i)
		k.setUpstream(upstream)
		upstream = k
	}
	for _, move := range moves {
		for i := 0; i < move.distance; i++ {
			head.move(move.direction)
		}
	}

	total := 0
	for _, m := range upstream.seen {
		total += len(m)
	}
	return total
}
