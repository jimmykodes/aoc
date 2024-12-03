package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	p1()
	p2()
}

func p2() {
	data, err := os.ReadFile("assets/input.txt")
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	lines := bytes.Split(data, []byte("\n"))
	var l1 []int
	l2 := map[int]int{}
	for _, line := range lines {
		p1, p2, _ := strings.Cut(string(line), "   ")
		l1 = append(l1, mustparse(p1))
		l2[mustparse(p2)] += 1
	}
	var similarity int
	for _, i := range l1 {
		similarity += i * l2[i]
	}
	fmt.Println(similarity)
}

func p1() {
	data, err := os.ReadFile("assets/input.txt")
	if err != nil {
		panic(err)
	}
	data = bytes.TrimSpace(data)
	lines := bytes.Split(data, []byte("\n"))
	var l1, l2 []int
	for _, line := range lines {
		p1, p2, _ := strings.Cut(string(line), "   ")
		l1 = append(l1, mustparse(p1))
		l2 = append(l2, mustparse(p2))
	}
	slices.Sort(l1)
	slices.Sort(l2)
	var total int
	for i := range len(l1) {
		delta := l2[i] - l1[i]
		if delta < 0 {
			delta = -delta
		}
		total += delta
	}
	fmt.Println(total)
}

func mustparse(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
