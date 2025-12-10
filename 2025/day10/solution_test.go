package main

import (
	"fmt"
	"testing"
)

func Test_parseTarget(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"[.##.]", 0b0110},
		{"[...#.]", 0b01000},
		{"[#..#.]", 0b01001},
		{"[#.]", 0b01},
		{"[####]", 0b1111},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := parseTarget(tt.s)
			if got != tt.want {
				t.Errorf("parseTarget() = %b, want %b", got, tt.want)
			}
		})
	}
}

func Test_parseButton(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"(0)", 0b1},
		{"(0,1)", 0b11},
		{"(0,1,2)", 0b111},
		{"(0,2)", 0b101},
		{"(1,5)", 0b100010},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := parseButton(tt.s)
			if got != tt.want {
				t.Errorf("parseButton() = %v, want %v", got, tt.want)
			}
		})
	}
}
