package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	hands := getHands("assets/input.txt")
	// fmt.Println(p1(hands))
	fmt.Println(p2(hands))
}

func p1(hands []*Hand) int {
	slices.SortFunc(hands, func(a, b *Hand) int {
		if a.strength == 0 {
			a.strength = classify(a.counts)
		}
		if b.strength == 0 {
			b.strength = classify(b.counts)
		}
		// -1 when a < b; 1 when a > b; 0 when a == b
		if a.strength < b.strength {
			return -1
		}
		if a.strength > b.strength {
			return 1
		}
		// strengths are equal, check cards
		for i := 0; i < 5; i++ {
			if a.cards[i] < b.cards[i] {
				return -1
			}
			if a.cards[i] > b.cards[i] {
				return 1
			}
		}
		return 0
	})
	total := 0
	for rank, hand := range hands {
		total += hand.bid * (rank + 1)
	}
	return total
}

func p2(hands []*Hand) int {
	slices.SortFunc(hands, func(a, b *Hand) int {
		if a.strength == 0 {
			a.strength = classifyWild(a.counts)
		}
		if b.strength == 0 {
			b.strength = classifyWild(b.counts)
		}
		// -1 when a < b; 1 when a > b; 0 when a == b
		if a.strength < b.strength {
			return -1
		}
		if a.strength > b.strength {
			return 1
		}
		// strengths are equal, check cards
		for i := 0; i < 5; i++ {
			if a.cards[i] < b.cards[i] {
				return -1
			}
			if a.cards[i] > b.cards[i] {
				return 1
			}
		}
		return 0
	})
	total := 0
	for rank, hand := range hands {
		total += hand.bid * (rank + 1)
	}
	return total
}

type Hand struct {
	cards    [5]int
	counts   map[rune]int
	strength int
	bid      int
}

const (
	HighCard = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var cardStrength = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	// 'J': 11, // uncomment for part 1
	'J': 1, // comment out for part 1
	'Q': 12,
	'K': 13,
	'A': 14,
}

func getHands(fn string) []*Hand {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var out []*Hand
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		handStr, bid, _ := strings.Cut(line, " ")
		var hand Hand
		hand.bid, _ = strconv.Atoi(bid)

		hand.counts = make(map[rune]int, 5)
		for i, r := range handStr {
			hand.cards[i] = cardStrength[r]
			hand.counts[r]++
		}

		out = append(out, &hand)
	}
	return out
}

func classifyWild(hand map[rune]int) int {
	h := make([]int, 0, len(hand))
	wild := 0
	for r, v := range hand {
		if r == 'J' {
			wild = v
			continue
		}
		h = append(h, v)
	}
	slices.Sort(h)
	slices.Reverse(h)

	if wild == 5 || wild == 4 || h[0] == 5 {
		// the wild==4 case is because four wild cards would
		// match with the one remaining card to be five of a kind
		return FiveOfAKind
	}
	if h[0] == 4 {
		if wild == 1 {
			return FiveOfAKind
		}
		return FourOfAKind
	}
	if h[0] == 3 {
		if wild == 2 {
			return FiveOfAKind
		}
		if wild == 1 {
			return FourOfAKind
		}
		if h[1] == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	}
	if h[0] == 2 {
		if wild == 3 {
			return FiveOfAKind
		}
		if wild == 2 {
			return FourOfAKind
		}
		if h[1] == 2 {
			if wild == 1 {
				return FullHouse
			}
			return TwoPair
		}
		if wild == 1 {
			return ThreeOfAKind
		}
		return OnePair
	}
	if wild == 3 {
		return FourOfAKind
	}
	if wild == 2 {
		return ThreeOfAKind
	}
	if wild == 1 {
		return OnePair
	}
	return HighCard
}

func classify(hand map[rune]int) int {
	r := make([]int, 0, len(hand))
	for _, v := range hand {
		r = append(r, v)
	}
	slices.Sort(r)
	slices.Reverse(r)

	if r[0] == 5 {
		return FiveOfAKind
	}
	if r[0] == 4 {
		return FourOfAKind
	}
	if r[0] == 3 {
		if r[1] == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	}
	if r[0] == 2 {
		if r[1] == 2 {
			return TwoPair
		}
		return OnePair
	}
	return HighCard
}
