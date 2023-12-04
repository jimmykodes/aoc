package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	c, err := cards("assets/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(p2(c))
	start := time.Now()
	for i := 0; i < 10000; i++ {
		p2(c)
	}
	fmt.Println(time.Since(start).Seconds())
}

func p1(cards []*Card) int {
	total := 0.0
	for _, c := range cards {
		if c.wins > 0 {
			total += math.Pow(2, float64(c.wins-1))
		}
	}
	return int(total)
}

func p2(cards []*Card) int {
	cache := Cache{
		cache: make([]*int, len(cards)+1),
	}
	total := len(cards)
	for _, card := range cards {
		total += cache.wins(card, cards)
	}
	return total
}

type Card struct {
	id   int
	wins int
}

type Cache struct {
	cache []*int
}

func (c *Cache) wins(card *Card, cards []*Card) int {
	if wins := c.cache[card.id]; wins != nil {
		return *wins
	}
	total := card.wins
	for i := 0; i < card.wins; i++ {
		if card.id+i >= len(cards) {
			break
		}
		total += c.wins(cards[card.id+i], cards)
	}
	c.cache[card.id] = &total
	return total
}

func cards(filename string) ([]*Card, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var out []*Card
	for scanner.Scan() {
		line := scanner.Text()
		card, nums, _ := strings.Cut(line, ": ")
		fields := strings.Fields(card)
		var c Card
		c.id, _ = strconv.Atoi(fields[1])

		winningStr, haveStr, _ := strings.Cut(nums, " | ")

		winning := strToInts(winningStr)
		have := strToInts(haveStr)
		for k := range have {
			if _, ok := winning[k]; ok {
				c.wins++
			}
		}
		out = append(out, &c)
	}
	return out, nil
}

func strToInts(s string) map[int]struct{} {
	f := strings.Fields(s)
	out := make(map[int]struct{}, len(f))
	for _, field := range f {
		k, _ := strconv.Atoi(field)
		out[k] = struct{}{}
	}
	return out
}
