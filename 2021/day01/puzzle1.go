package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	attempt1()
	fmt.Println("attempt1 completed in", time.Since(start))
	start = time.Now()
	attempt2()
	fmt.Println("attempt2 completed in", time.Since(start))
}

func attempt1() {
	data, err := os.ReadFile("assets/input.txt")
	if err != nil {
		panic(err)
	}
	rows := bytes.Split(data, []byte("\n"))
	var (
		last  int
		count int
	)
	for i, row := range rows {
		if len(row) == 0 {
			continue
		}
		current, err := strconv.Atoi(string(row))
		if err != nil {
			panic(err)
		}
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

func attempt2() {
	f, err := os.Open("assets/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	last := 0
	count := -1
	for scanner.Scan() {
		current, _ := strconv.Atoi(scanner.Text())
		if current > last {
			count++
		}
		last = current
	}
	fmt.Println(count)
}
