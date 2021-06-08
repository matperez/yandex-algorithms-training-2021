package main

import (
	"fmt"
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		k1 int64
		m  int64
		k2 int64
		n2 int64
		p2 int64
		p1 int64
		n1 int64
	}{
		{
			k1: 89,
			m:  20,
			k2: 41,
			p2: 1,
			n2: 11,
			p1: 2,
			n1: 3,
		},
		{
			k1: 11,
			m:  1,
			k2: 1,
			p2: 1,
			n2: 1,
			p1: 0,
			n1: 1,
		},
		{
			k1: 3,
			m:  2,
			k2: 2,
			p2: 2,
			n2: 1,
			p1: -1,
			n1: -1,
		},
		{
			k1: 5,
			m:  20,
			k2: 2,
			p2: 1,
			n2: 1,
			p1: 1,
			n1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d %d %d %d", tt.k1, tt.m, tt.k2, tt.p2, tt.n2), func(t *testing.T) {
			p1, n1 := solve(tt.k1, tt.m, tt.k2, tt.p2, tt.n2)
			if p1 != tt.p1 {
				t.Errorf("got entrance = %v, want %v", p1, tt.p1)
			}
			if n1 != tt.n1 {
				t.Errorf("got floor = %v, want %v", n1, tt.n1)
			}
		})
	}
}

func Test_valid(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		m    int64
		k2   int64
		p2   int64
		n2   int64
		want bool
	}{
		{
			m:    20,
			k2:   41,
			p2:   1,
			n2:   11,
			want: true,
		},
		{
			m:  1,
			k2: 1,
			p2: 1,
			n2: 1,
			want: true,
		},
		{
			m:  2,
			k2: 2,
			p2: 2,
			n2: 1,
			want: false,
		},
		{
			m:  20,
			k2: 2,
			p2: 1,
			n2: 1,
			want: true,
		},
		{
			m:  9,
			k2: 62,
			p2: 2,
			n2: 6,
			want: true,
		},
		{
			m:  9,
			k2: 61,
			p2: 2,
			n2: 6,
			want: true,
		},
		{
			m:  9,
			k2: 60,
			p2: 2,
			n2: 6,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d %d %d", tt.m, tt.k2, tt.p2, tt.n2), func(t *testing.T) {
			if got := valid(tt.m, tt.k2, tt.p2, tt.n2); got != tt.want {
				t.Errorf("valid() = %v, want %v", got, tt.want)
			}
		})
	}
}