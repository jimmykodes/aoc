package main

type direction int

func (d direction) String() string {
	return string("KRLUD"[d])
}

const (
	unknown direction = iota
	right
	left
	up
	down
)

func directionFromStr(s string) direction {
	switch s {
	case "R":
		return right
	case "L":
		return left
	case "U":
		return up
	case "D":
		return down
	default:
		return unknown
	}
}
