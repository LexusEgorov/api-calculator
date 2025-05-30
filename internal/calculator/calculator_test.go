package calculator

import "testing"

type TestCase struct {
	Input    []float64
	Expected float64
}

func TestSumNums(t *testing.T) {
	tests := []TestCase{
		{
			Input:    []float64{1.2, 5.3, 2.2},
			Expected: 8.7,
		},
		{
			Input:    []float64{10000000, 20000000, 30000000},
			Expected: 60000000,
		},
		{
			Input:    []float64{-1, -51},
			Expected: -52,
		},
		{
			Input:    []float64{0, 0, 0},
			Expected: 0,
		},
		{
			Input:    []float64{},
			Expected: 0,
		},
	}

	for i, test := range tests {
		res := sumNums(test.Input...)

		if res != test.Expected {
			t.Errorf("case #%d: not equal! got: %f expected %f", i+1, res, test.Expected)
		}
	}
}

func TestMultNums(t *testing.T) {
	tests := []TestCase{
		{
			Input:    []float64{1.5, 3, 2},
			Expected: 9,
		},
		{
			Input:    []float64{10000000, 20000000, 0},
			Expected: 0,
		},
		{
			Input:    []float64{-1, -51},
			Expected: 51,
		},
		{
			Input:    []float64{0, 0, 0},
			Expected: 0,
		},
		{
			Input:    []float64{},
			Expected: 0,
		},
	}

	for i, test := range tests {
		res := multNums(test.Input...)

		if res != test.Expected {
			t.Errorf("case #%d: not equal! got: %f expected %f", i+1, res, test.Expected)
		}
	}
}
