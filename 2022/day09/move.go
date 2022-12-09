package main

import (
	"fmt"
	"strconv"
	"strings"
)

func newMove(line string) (*move, error) {
	l, r, _ := strings.Cut(line, " ")
	dir := directionFromStr(l)
	dist, err := strconv.Atoi(r)
	if err != nil {
		return nil, err
	}
	return &move{
		direction: dir,
		distance:  dist,
	}, nil
}

type move struct {
	direction direction
	distance  int
}

func (m *move) String() string {
	return fmt.Sprintf("%v - %d", m.direction, m.distance)
}
