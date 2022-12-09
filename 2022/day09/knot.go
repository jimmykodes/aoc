package main

import "fmt"

type record map[int]map[int]bool

func (r record) add(c coord) {
	m, ok := r[c.x]
	if !ok {
		m = make(map[int]bool)
		r[c.x] = m
	}
	m[c.y] = true
}

func newKnot(id int) *knot {
	k := &knot{
		id:       id,
		position: coord{},
		seen:     make(record),
	}
	k.seen.add(k.position)
	return k
}

type knot struct {
	id         int
	position   coord
	upstream   *knot
	downstream *knot
	seen       record
}

func (k *knot) setUpstream(upstream *knot) {
	k.upstream = upstream
	upstream.downstream = k
}

func (e *knot) move(dirs ...direction) {
	for _, dir := range dirs {
		switch dir {
		case right:
			e.position.x++
		case left:
			e.position.x--
		case up:
			e.position.y++
		case down:
			e.position.y--
		}
	}
	if e.downstream != nil {
		e.downstream.update()
	}
}

func (e *knot) update() {
	if e.upstream == nil {
		fmt.Println("found an unconnected end")
		return
	}
	if e.position.touching(e.upstream.position) {
		return
	}
	e.move(e.position.direction(e.upstream.position)...)
	e.seen.add(e.position)
	if !e.position.touching(e.upstream.position) {
		fmt.Println("something went wrong, things aren't touching", e.position, e.upstream.position)
	}
}
