package main

import "testing"

func Test_stripNonDigits(t *testing.T) {
	tests := []struct {
		num string
		want string
	}{
		{
			num: "8(495)430-23-97",
			want: "84954302397",
		},
		{
			num: "+7-4-9-5-43-023-97",
			want: "74954302397",
		},
		{
			num: "4-3-0-2-3-9-7",
			want: "4302397",
		},
	}
	for _, tt := range tests {
		t.Run(tt.num, func(t *testing.T) {
			if got := stripNonDigits(tt.num); got != tt.want {
				t.Errorf("stripNonDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendCode(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		num string
		want string
	}{
		{
			num: "4302397",
			want: "4954302397",
		},
		{
			num: "74954302397",
			want: "74954302397",
		},
	}
	for _, tt := range tests {
		t.Run(tt.num, func(t *testing.T) {
			if got := appendCode(tt.num); got != tt.want {
				t.Errorf("appendCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stripCountryCode(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		num string
		want string
	}{
		{
			num: "74954302397",
			want: "4954302397",
		},
		{
			num: "84954302397",
			want: "4954302397",
		},
		{
			num: "4302397",
			want: "4302397",
		},
	}
	for _, tt := range tests {
		t.Run(tt.num, func(t *testing.T) {
			if got := stripCountryCode(tt.num); got != tt.want {
				t.Errorf("stripCountryCode() = %v, want %v", got, tt.want)
			}
		})
	}
}