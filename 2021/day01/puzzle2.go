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
	rows := bytes.Split(data, []byte("\n"))
	var (
		last  int
		count int
	)
	for i := 0; i < len(rows)-2; i++ {
		current := sum(rows[i], rows[i+1], rows[i+2])
		fmt.Println("considering indexes: ", i, i+1, i+2, "with values", string(rows[i]), string(rows[i+1]), string(rows[i+2]), "and sum", current)
		if i == 0 {
			last = current
			continue
		}
		if current > last {
			count++
		}
		last = current
	}
	fmt.Println(count)
}

func sum(inputs ...[]byte) int {
	var s int
	for _, b := range inputs {
		str := string(b)
		if str == "" {
			continue
		}
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		s += i
	}
	return s
}
