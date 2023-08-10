package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
)

var (
	//go:embed test.bkup.txt
	test []byte
	//go:embed input.txt
	input []byte

	data = input
)

type status int

const (
	unknown status = iota
	ordered
	unordered
)

func comp(left, right any) status {
	switch l := left.(type) {
	case []any:
		switch r := right.(type) {
		case []any:
			for j := 0; j < len(l); j++ {
				if j >= len(r) {
					return unordered
				}
				if s := comp(l[j], r[j]); s != unknown {
					return s
				}
			}
			if len(l) < len(r) {
				return ordered
			}
			return unknown
		case float64:
			return comp(left, []any{r})
		default:
			panic("invalid type")
		}
	case float64:
		switch r := right.(type) {
		case []any:
			return comp([]any{l}, r)
		case float64:
			if l != r {
				if l <= r {
					return ordered
				}
				return unordered
			}
			return unknown
		default:
			panic("invalid type")
		}
	default:
		panic("invalid type")
	}
}

func getData() []any {
	var (
		lines = bytes.Split(data, []byte("\n"))
		out   []any
	)
	for i := 2; i < len(lines)+1; i += 3 {
		first := lines[i-2]
		second := lines[i-1]
		var l, r []any
		if err := json.Unmarshal(first, &l); err != nil {
			panic(err)
		}
		if err := json.Unmarshal(second, &r); err != nil {
			panic(err)
		}
		out = append(out, l, r)
	}
	return out
}

func p1() int {
	packets := getData()
	total := 0
	for i := 1; i < len(packets); i += 2 {
		if comp(packets[i-1], packets[i]) != unordered {
			total += i/2 + 1
		}
	}
	return total
}
func p2() int {
	packets := append(getData(), []any{[]any{2.}}, []any{[]any{6.}})
	sort.Slice(packets, func(i, j int) bool { return comp(packets[i], packets[j]) != unordered })
	key := 1
	for i, pkt := range packets {
		if fmt.Sprint(pkt) == "[[2]]" || fmt.Sprint(pkt) == "[[6]]" {
			key *= i + 1
		}
	}
	return key
}

func main() {
	fmt.Println(p1())
	fmt.Println(p2())
}
