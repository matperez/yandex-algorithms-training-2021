package main

import (
	"fmt"
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		m    int
		want int
	}{
		{
			n:    10,
			k:    5,
			m:    2,
			want: 4,
		},
		{
			n:    13,
			k:    5,
			m:    3,
			want: 3,
		},
		{
			n:    14,
			k:    5,
			m:    3,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d %d", tt.n, tt.k, tt.m), func(t *testing.T) {
			if got := solve(tt.n, tt.k, tt.m); got != tt.want {
				t.Errorf("solve() got %v, want %v", got, tt.want)
			}
		})
	}
}
