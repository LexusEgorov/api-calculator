package calculator

import (
	"testing"
)

func Test_sumNums(t *testing.T) {
	type args struct {
		nums []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "regular case",
			args: args{
				nums: []float64{1, 2, 3},
			},
			want: float64(6),
		},
		{
			name: "zeros",
			args: args{
				nums: []float64{0, 0, 0},
			},
			want: float64(0),
		},
		{
			name: "negative nums",
			args: args{
				nums: []float64{-1, -2, -3},
			},
			want: float64(-6),
		},
		{
			name: "empty input",
			args: args{
				nums: []float64{},
			},
			want: float64(0),
		},
		{
			name: "big input",
			args: args{
				nums: []float64{1000000000, 2000000000, 3000000000},
			},
			want: float64(6000000000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNums(tt.args.nums...); got != tt.want {
				t.Errorf("sumNums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multNums(t *testing.T) {
	type args struct {
		nums []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "regular case",
			args: args{
				nums: []float64{1, 2, 3},
			},
			want: float64(6),
		},
		{
			name: "zeros",
			args: args{
				nums: []float64{0, 0, 0},
			},
			want: float64(0),
		},
		{
			name: "negative nums",
			args: args{
				nums: []float64{-1, -2, -3},
			},
			want: float64(-6),
		},
		{
			name: "empty input",
			args: args{
				nums: []float64{},
			},
			want: float64(0),
		},
		{
			name: "big input",
			args: args{
				nums: []float64{1000000000, 2000000000, 3000000000},
			},
			want: float64(6000000000000000000000000000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multNums(tt.args.nums...); got != tt.want {
				t.Errorf("multNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
