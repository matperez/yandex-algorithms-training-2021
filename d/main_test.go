package main

import (
	"strconv"
	"testing"
)

func Test_check(t *testing.T) {
	tests := []struct {
		x int64
		a int64
		b int64
		c int64
		want bool
	}{
		{
			x:    0,
			a:    1,
			b:    0,
			c:    0,
			want: true,
		},
		{
			x:    7,
			a:    1,
			b:    2,
			c:    3,
			want: true,
		},
		{
			x:    7,
			a:    1,
			b:    2,
			c:    -3,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(strconv.FormatInt(tt.c, 10), func(t *testing.T) {
			if got := check(tt.x, tt.a, tt.b, tt.c); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
