package main

import (
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
)

var filename = "assets/input.txt"

func main() {
	p1()

	p2()
}

func p1() {
	total := 0
	for o, strs := range parse(filename) {
		rowTotal := 0
		var combine func(a, b int) int
		switch o {
		case '+':
			combine = func(a, b int) int { return a + b }
		case '*':
			rowTotal = 1
			combine = func(a, b int) int { return a * b }
		default:
			panic("invalid op")
		}
		for _, str := range strs {
			i, err := strconv.Atoi(strings.TrimSpace(str))
			if err != nil {
				panic(err)
			}
			rowTotal = combine(rowTotal, i)
		}
		total += rowTotal
	}
	fmt.Println(total)
}

func p2() {
	total := 0
	for o, strs := range parse(filename) {
		rowTotal := 0
		var combine func(a, b int) int
		switch o {
		case '+':
			combine = func(a, b int) int { return a + b }
		case '*':
			rowTotal = 1
			combine = func(a, b int) int { return a * b }
		default:
			panic("invalid op")
		}
		strs = transform(strs)
		for _, str := range strs {
			val := strings.TrimSpace(str)
			if val == "" {
				continue
			}
			i, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			rowTotal = combine(rowTotal, i)
		}
		total += rowTotal
	}
	fmt.Println(total)
}

func transform(strs []string) []string {
	var out []string
	for i := range strs[0] {
		var d []byte
		for _, str := range strs {
			d = append(d, str[i])
		}
		out = append(out, string(d))
	}

	return out
}

func parse(fname string) iter.Seq2[byte, []string] {
	data, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	rows, opRowStrs := rows[:len(rows)-1], rows[len(rows)-1]
	opIdx := 0
	return func(yield func(byte, []string) bool) {
		for {
			nextOpIdx := opIdx + 1
			for nextOpIdx < len(opRowStrs) && opRowStrs[nextOpIdx] == ' ' {
				nextOpIdx++
			}
			rowData := make([]string, len(rows))

			for i, row := range rows {
				rowData[i] = row[opIdx:nextOpIdx]
			}
			if !yield(opRowStrs[opIdx], rowData) {
				return
			}
			if nextOpIdx == len(opRowStrs) {
				return
			}
			opIdx = nextOpIdx
		}
	}
}
