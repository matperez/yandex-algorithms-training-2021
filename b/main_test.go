package main

import "testing"

func Test_check(t *testing.T) {
	type args struct {
		a int64
		b int64
		c int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: struct {
				a int64
				b int64
				c int64
			}{a: 3, b: 4, c: 5},
			want: true,
		},
		{
			name: "1",
			args: struct {
				a int64
				b int64
				c int64
			}{a: 3, b: 5, c: 4},
			want: true,
		},
		{
			name: "1",
			args: struct {
				a int64
				b int64
				c int64
			}{a: 3, b: 5, c: 3},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
