package main

import (
	"reflect"
	"testing"
)

func Test_conslidateRanges(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ranges []Range
		want   []Range
	}{
		{
			name:   "no overlap",
			ranges: []Range{{0, 10}, {12, 15}},
			want:   []Range{{0, 10}, {12, 15}},
		},
		{
			name:   "completely contained",
			ranges: []Range{{0, 10}, {5, 8}},
			want:   []Range{{0, 10}},
		},
		{
			name: "failing",
			ranges: []Range{
				{31_901_931_864_283, 39_948_040_427_492},
				{34_688_333_353_698, 35_776_359_268_938},
			},
			want: []Range{
				{31_901_931_864_283, 39_948_040_427_492},
			},
		},

		{
			name:   "completely equal",
			ranges: []Range{{0, 10}, {0, 10}},
			want:   []Range{{0, 10}},
		},
		{
			name:   "3 span",
			ranges: []Range{{0, 10}, {8, 15}, {14, 20}},
			want:   []Range{{0, 20}},
		},
		{
			name:   "3 span out of order",
			ranges: []Range{{0, 10}, {14, 20}, {8, 15}},
			want:   []Range{{0, 15}, {14, 20}},
		},
		{
			name:   "touching ranges",
			ranges: []Range{{0, 5}, {5, 10}},
			want:   []Range{{0, 10}},
		},
		{
			name:   "single range",
			ranges: []Range{{5, 10}},
			want:   []Range{{5, 10}},
		},
		{
			name:   "partial overlap front",
			ranges: []Range{{0, 10}, {5, 15}},
			want:   []Range{{0, 15}},
		},
		{
			name:   "partial overlap back",
			ranges: []Range{{5, 15}, {0, 10}},
			want:   []Range{{0, 15}},
		},
		{
			name:   "multiple separate ranges",
			ranges: []Range{{0, 5}, {10, 15}, {20, 25}},
			want:   []Range{{0, 5}, {10, 15}, {20, 25}},
		},
		{
			name:   "complex multiple overlaps",
			ranges: []Range{{0, 5}, {3, 8}, {15, 20}, {18, 25}, {30, 35}},
			want:   []Range{{0, 8}, {15, 25}, {30, 35}},
		},
		{
			name:   "adjacent ranges",
			ranges: []Range{{0, 5}, {6, 11}},
			want:   []Range{{0, 11}},
		},
		{
			name:   "adjacent ranges reversed",
			ranges: []Range{{6, 11}, {0, 5}},
			want:   []Range{{0, 11}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := conslidateRanges(tt.ranges)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("conslidateRanges() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_completelyConslidate(t *testing.T) {
	tests := []struct {
		name   string
		ranges []Range
		want   []Range
	}{
		{
			name:   "no consolidation needed",
			ranges: []Range{{0, 5}, {10, 15}, {20, 25}},
			want:   []Range{{0, 5}, {10, 15}, {20, 25}},
		},
		{
			name:   "single pass consolidation",
			ranges: []Range{{0, 5}, {3, 8}},
			want:   []Range{{0, 8}},
		},
		{
			name:   "3 span out of order",
			ranges: []Range{{0, 10}, {14, 20}, {8, 15}},
			want:   []Range{{0, 20}},
		},
		{
			name:   "multi-pass consolidation",
			ranges: []Range{{0, 5}, {3, 8}, {7, 12}, {10, 15}},
			want:   []Range{{0, 15}},
		},
		{
			name:   "complex chain consolidation",
			ranges: []Range{{0, 5}, {10, 15}, {3, 12}, {14, 20}, {18, 25}},
			want:   []Range{{0, 25}},
		},
		{
			name:   "multiple separate chains",
			ranges: []Range{{0, 5}, {3, 8}, {20, 25}, {23, 30}, {40, 45}},
			want:   []Range{{0, 8}, {20, 30}, {40, 45}},
		},
		{
			name:   "all ranges merge into one",
			ranges: []Range{{0, 10}, {5, 15}, {12, 20}, {18, 25}, {22, 30}},
			want:   []Range{{0, 30}},
		},
		{
			name:   "single range",
			ranges: []Range{{5, 10}},
			want:   []Range{{5, 10}},
		},
		{
			name:   "touching ranges that should merge",
			ranges: []Range{{0, 5}, {5, 10}, {10, 15}},
			want:   []Range{{0, 15}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := completelyConslidate(tt.ranges)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("completelyConslidate() = %v, want %v", got, tt.want)
			}
		})
	}
}
