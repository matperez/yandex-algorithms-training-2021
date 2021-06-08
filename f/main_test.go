package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		a1   int
		b1   int
		a2   int
		b2   int
		want [][]int
	}{
		{
			a1: 10,
			b1: 2,
			a2: 2,
			b2: 10,
			want: [][]int{
				{20, 2},
				{2, 20},
				{4, 10},
				{10, 4},
			},
		},
		{
			a1: 5,
			b1: 7,
			a2: 3,
			b2: 2,
			want: [][]int{
				{9, 5},
				{5, 9},
			},
		},
		{
			a1: 3,
			b1: 5,
			a2: 6,
			b2: 4,
			want: [][]int{
				{7, 6},
				{6, 7},
			},
		},
		{
			a1: 1000,
			b1: 1000,
			a2: 1000,
			b2: 900,
			want: [][]int{
				{1000, 1900},
				{1900, 1000},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d %d %d", tt.a1, tt.b1, tt.a2, tt.b2), func(t *testing.T) {
			x, y := solve(tt.a1, tt.b1, tt.a2, tt.b2)
			got := []int{x, y}
			for _, want := range tt.want {
				if reflect.DeepEqual(got, want) {
					return
				}
			}
			t.Errorf("solve() got [%v %v], want %v", x, y, tt.want)
		})
	}
}
