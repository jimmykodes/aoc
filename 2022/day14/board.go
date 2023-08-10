package main

import "strings"

type State int

const (
	Air State = iota
	Wall
	Sand
)

func NewBoard(width, height int) *Board {
	b := &Board{}
	b.board = make([][]State, height+2)
	for i := 0; i < len(b.board); i++ {
		b.board[i] = make([]State, width+1000)
	}
	return b
}

type Board struct {
	board [][]State
}

func (b *Board) String() string {
	var sb strings.Builder
	for _, row := range b.board {
		for _, col := range row {
			switch col {
			case Air:
				sb.WriteString(".")
			case Wall:
				sb.WriteString("#")
			default:
				sb.WriteString("o")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
