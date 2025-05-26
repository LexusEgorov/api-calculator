package clevercalc

import (
	"reflect"
	"testing"

	"github.com/LexusEgorov/api-calculator/internal/models"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()
var calculator = New(logger)

func TestNew(t *testing.T) {
	calc := Calculator{
		priorityMap: newPriority(),
		parser:      newParser(),
		logger:      logger,
	}

	if !reflect.DeepEqual(calc.priorityMap, calculator.priorityMap) {
		t.Errorf("not equal: got '%v', expected '%v'", calc.priorityMap, calculator.priorityMap)
	}

	if !reflect.DeepEqual(calc.parser, calculator.parser) {
		t.Errorf("not equal: got '%v', expected '%v'", calc.parser, calculator.parser)
	}

	if !reflect.DeepEqual(calc.logger, calculator.logger) {
		t.Errorf("not equal: got '%v', expected '%v'", calc.logger, calculator.logger)
	}
}

func TestCompute(t *testing.T) {
	var expected float64 = 4
	input := []string{
		"2",
		"2",
		"+",
	}

	res, err := calculator.compute(input)

	if err != nil {
		t.Errorf("got error: %v", err)
	}

	if res != expected {
		t.Errorf("not equal: got '%v', expected '%v'", res, expected)
	}
}

func TestComputeErr(t *testing.T) {
	input := []string{
		"2",
		"2",
	}

	_, err := calculator.compute(input)

	if err == nil {
		t.Error("expected error")
	}
}

func TestCalcCompute(t *testing.T) {
	var expected float64 = 4
	input := "2+2"

	res, err := calculator.Compute(input)

	if err != nil {
		t.Errorf("got error: %v", err)
	}

	if res != expected {
		t.Errorf("not equal: got '%v', expected '%v'", res, expected)
	}
}

func TestCalcComputeErr(t *testing.T) {
	input := "(2+2"

	_, err := calculator.Compute(input)

	if err == nil {
		t.Error("expected error")
	}
}

type Input struct {
	a         float64
	b         float64
	operation string
}

type Test struct {
	Input    Input
	Expected float64
}

func TestComputeFn(t *testing.T) {
	tests := []Test{
		{
			Input: Input{
				a:         4,
				b:         2,
				operation: models.OPERATION_DIV,
			},
			Expected: 2,
		},
		{
			Input: Input{
				a:         4,
				b:         2,
				operation: models.OPERATION_POW,
			},
			Expected: 16,
		},
		{
			Input: Input{
				a:         4,
				b:         2,
				operation: models.OPERATION_SUM,
			},
			Expected: 6,
		},
		{
			Input: Input{
				a:         4,
				b:         2,
				operation: models.OPERATION_SUB,
			},
			Expected: 2,
		},
		{
			Input: Input{
				a:         4,
				b:         2,
				operation: models.OPERATION_MULT,
			},
			Expected: 8,
		},
	}

	for i, test := range tests {
		res, err := compute(test.Input.a, test.Input.b, test.Input.operation)

		if err != nil {
			t.Errorf("test case #%d: got error: %v", i, err)
		}

		if res != test.Expected {
			t.Errorf("test case #%d: got %v, expected %v", i, res, test.Expected)
		}
	}
}

func TestComputeFnError(t *testing.T) {
	_, err := compute(2, 0, models.OPERATION_DIV)

	if err == nil {
		t.Error("expected error")
	}
}

func TestGetNum(t *testing.T) {
	input := "10"
	var expected float64 = 10
	stack := Stack{}
	stack.Push(input)

	res, err := getNum(&stack)

	if err != nil {
		t.Errorf("got error: %v", err)
	}

	if res != expected {
		t.Errorf("got '%v', expected '%v'", res, expected)
	}
}
