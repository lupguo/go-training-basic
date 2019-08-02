package ufun

import (
	"testing"
)

func TestAbs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "T1", args: args{1}, want: 1},
		{name: "T2", args: args{2}, want: 2},
		{name: "T3", args: args{-1}, want: 1},
		{name: "T4", args: args{-2}, want: 2},
		{name: "T5", args: args{0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.n); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
