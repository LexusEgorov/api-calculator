package clevercalc

import (
	"testing"

	"github.com/LexusEgorov/api-calculator/internal/models"
	"github.com/sirupsen/logrus"
)

func Test_getNum(t *testing.T) {
	testStack := Stack{}

	type args struct {
		from *Stack
	}
	tests := []struct {
		name    string
		args    args
		init    []string
		want    float64
		wantErr bool
	}{
		{
			name: "regular case",
			args: args{
				from: &testStack,
			},
			init: []string{
				"2",
				"2",
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "negative number",
			args: args{
				from: &testStack,
			},
			init: []string{
				"2",
				"-2",
			},
			want:    -2,
			wantErr: false,
		},
		{
			name: "float number",
			args: args{
				from: &testStack,
			},
			init: []string{
				"2",
				"2.2",
			},
			want:    2.2,
			wantErr: false,
		},
		{
			name: "big number",
			args: args{
				from: &testStack,
			},
			init: []string{
				"2",
				"88888888888888",
			},
			want:    88888888888888,
			wantErr: false,
		},
		{
			name: "small number",
			args: args{
				from: &testStack,
			},
			init: []string{
				"2",
				"0.00000000005",
			},
			want:    0.00000000005,
			wantErr: false,
		},
		{
			name: "empty stack",
			args: args{
				from: &testStack,
			},
			init:    []string{},
			want:    0,
			wantErr: true,
		},
		{
			name: "not number in stack",
			args: args{
				from: &testStack,
			},
			init: []string{
				"a",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.from.data = tt.init
			tt.args.from.size = len(tt.init)

			got, err := getNum(tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("getNum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculator_compute(t *testing.T) {
	loggerTest := logrus.New()
	calculatorTest := New(loggerTest)

	type args struct {
		parsed []string
	}
	tests := []struct {
		name    string
		c       Calculator
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "simple",
			c:    *calculatorTest,
			args: args{
				parsed: []string{
					"2",
					"2",
					"+",
				},
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "many actions",
			c:    *calculatorTest,
			args: args{
				parsed: []string{
					"8",
					"4",
					"3",
					"3",
					"/",
					"+",
					"8",
					"/",
					"*",
					"15",
					"+",
				},
			},
			want:    20,
			wantErr: false,
		},
		{
			name: "negative num",
			c:    *calculatorTest,
			args: args{
				parsed: []string{
					"2",
					"-",
				},
			},
			want:    -2,
			wantErr: false,
		},
		{
			name: "operations problem",
			c:    *calculatorTest,
			args: args{
				parsed: []string{
					"*",
					"2",
					"-",
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "div by zero",
			c:    *calculatorTest,
			args: args{
				parsed: []string{
					"5",
					"0",
					"/",
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "unknown action",
			c:    *calculatorTest,
			args: args{
				parsed: []string{
					"5",
					"0",
					"sin",
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.compute(tt.args.parsed)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculator.compute(\"%v\") error = %v, wantErr %v", tt.args.parsed, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculator.compute(\"%v\") = %v, want %v", tt.args.parsed, got, tt.want)
			}
		})
	}
}

func Test_compute(t *testing.T) {
	type args struct {
		a         float64
		b         float64
		operation string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "sum",
			args: args{
				a:         20,
				b:         4,
				operation: models.OperationSum,
			},
			want:    24,
			wantErr: false,
		},
		{
			name: "mult",
			args: args{
				a:         20,
				b:         4,
				operation: models.OperationMult,
			},
			want:    80,
			wantErr: false,
		},
		{
			name: "sub",
			args: args{
				a:         20,
				b:         4,
				operation: models.OperationSub,
			},
			want:    16,
			wantErr: false,
		},
		{
			name: "div",
			args: args{
				a:         20,
				b:         4,
				operation: models.OperationDiv,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "div on zero",
			args: args{
				a:         20,
				b:         0,
				operation: models.OperationDiv,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "pow",
			args: args{
				a:         2,
				b:         10,
				operation: models.OperationPow,
			},
			want:    1024,
			wantErr: false,
		},
		{
			name: "unknown",
			args: args{
				a:         2,
				b:         10,
				operation: "cos",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := compute(tt.args.a, tt.args.b, tt.args.operation)
			if (err != nil) != tt.wantErr {
				t.Errorf("compute(\"%v\", \"%v\", \"%v\") error = %v, wantErr %v", tt.args.a, tt.args.b, tt.args.operation, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("compute(\"%v\", \"%v\", \"%v\") = %v, want %v", tt.args.a, tt.args.b, tt.args.operation, got, tt.want)
			}
		})
	}
}

func TestCalculator_Compute(t *testing.T) {
	loggerTest := logrus.New()
	calculatorTest := New(loggerTest)

	type args struct {
		input string
	}
	tests := []struct {
		name    string
		c       Calculator
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "simple",
			c:    *calculatorTest,
			args: args{
				input: "2+2",
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "spaces",
			c:    *calculatorTest,
			args: args{
				input: "2             +                    2",
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "many acitons",
			c:    *calculatorTest,
			args: args{
				input: "8*((4+3/3)/8)+15-2^4",
			},
			want:    4,
			wantErr: false,
		},
		{
			name: "negative at start",
			c:    *calculatorTest,
			args: args{
				input: "-8*4+3",
			},
			want:    -29,
			wantErr: false,
		},
		{
			name: "negative after action",
			c:    *calculatorTest,
			args: args{
				input: "8*-4+3",
			},
			want:    -29,
			wantErr: false,
		},
		{
			name: "negative at brakes start",
			c:    *calculatorTest,
			args: args{
				input: "8*(-4+3)",
			},
			want:    -8,
			wantErr: false,
		},
		{
			name: "bad brakes #1",
			c:    *calculatorTest,
			args: args{
				input: "8*((4+3)",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "bad brakes #2",
			c:    *calculatorTest,
			args: args{
				input: "8*4+3)",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "div by zero",
			c:    *calculatorTest,
			args: args{
				input: "8/0",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Compute(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculator.Compute(\"%v\") error = %v, wantErr %v", tt.args.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculator.Compute(\"%v\") = %v, want %v", tt.args.input, got, tt.want)
			}
		})
	}
}
