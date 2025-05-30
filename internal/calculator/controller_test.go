package calculator

import (
	"testing"

	"github.com/LexusEgorov/api-calculator/internal/models"
	"github.com/sirupsen/logrus"
)

type mockCacher struct{}

func (m mockCacher) Get(input string, action models.Action) (models.CalcAction, error) {
	return models.CalcAction{}, models.ErrCacheNotFound
}

func (m mockCacher) Set(action models.CalcAction) error {
	return nil
}

type mockStorager struct{}

func (m mockStorager) Get(uID string) []models.CalcAction {
	return history
}

func (m mockStorager) Set(uID string, action models.CalcAction) {}

var history = make([]models.CalcAction, 0)
var cache = mockCacher{}
var storage = mockStorager{}
var controller = newController(logrus.New(), cache, storage)

type Test struct {
	Input    models.Input
	Expected models.CalcAction
}

func TestControllerSum(t *testing.T) {
	tests := []Test{
		{
			Input: models.Input{
				Input: "1, 1",
			},
			Expected: models.CalcAction{
				Input:  "1, 1",
				Action: models.Sum,
				Result: 2,
			},
		},
		{
			Input: models.Input{
				Input: "-5, -2",
			},
			Expected: models.CalcAction{
				Input:  "-5, -2",
				Action: models.Sum,
				Result: -7,
			},
		},
		{
			Input: models.Input{
				Input: "12345648, -8454515, 1884686",
			},
			Expected: models.CalcAction{
				Input:  "12345648, -8454515, 1884686",
				Action: models.Sum,
				Result: 5775819,
			},
		},
	}

	for i, test := range tests {
		res, err := controller.Sum("123", test.Input)

		if err != nil {
			t.Errorf("Got error: %v", err)
		}

		if *res != test.Expected {
			t.Errorf("Case #%d: Not equal!\nGot: %v\nExpected: %v\n", i+1, *res, test.Expected)
		}
	}
}

func TestControllerSumWithErr(t *testing.T) {
	input := models.Input{
		Input: "1, 1a",
	}
	_, err := controller.Sum("123", input)

	if err == nil {
		t.Error("Err nil, but expected not!")
	}
}

func TestControllerMult(t *testing.T) {
	tests := []Test{
		{
			Input: models.Input{
				Input: "1, 1",
			},
			Expected: models.CalcAction{
				Input:  "1, 1",
				Action: models.Mult,
				Result: 1,
			},
		},
		{
			Input: models.Input{
				Input: "-5, -2",
			},
			Expected: models.CalcAction{
				Input:  "-5, -2",
				Action: models.Mult,
				Result: 10,
			},
		},
		{
			Input: models.Input{
				Input: "5, 4, 0",
			},
			Expected: models.CalcAction{
				Input:  "5, 4, 0",
				Action: models.Mult,
				Result: 0,
			},
		},
	}

	for i, test := range tests {
		res, err := controller.Mult("123", test.Input)

		if err != nil {
			t.Errorf("Got error: %v", err)
		}

		if *res != test.Expected {
			t.Errorf("Case #%d: Not equal!\nGot: %v\nExpected: %v\n", i+1, *res, test.Expected)
		}
	}
}

func TestControllerMultWithErr(t *testing.T) {
	input := models.Input{
		Input: "1, 1a",
	}
	_, err := controller.Sum("123", input)

	if err == nil {
		t.Error("Err nil, but expected not!")
	}
}

func TestControllerCalculate(t *testing.T) {
	input := models.Input{
		Input: "(5-1)*(3+2)",
	}
	res, err := controller.Calculate("123", input)

	expected := models.CalcAction{
		Input:  input.Input,
		Action: models.Calc,
		Result: 20,
	}

	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if *res != expected {
		t.Errorf("Not equal!\nGot: %v\nExpected: %v\n", *res, expected)
	}
}

func TestControllerCalculateWithErr(t *testing.T) {
	input := models.Input{
		Input: "(1+2))",
	}
	_, err := controller.Calculate("123", input)

	if err == nil {
		t.Error("Err nil, but expected not!")
	}
}

type HistoryTest struct {
	history []models.CalcAction
}

func TestControllerHistory(t *testing.T) {
	tests := []HistoryTest{
		{
			history: make([]models.CalcAction, 0),
		},
		{
			history: []models.CalcAction{
				{
					Input:  "1, 1",
					Action: models.Sum,
					Result: 2,
				},
			},
		},
	}

	for i, test := range tests {
		history = test.history
		res := controller.History("")

		if !isEqual(res, test.history) {
			t.Errorf("test #%d: not equal! got: %v expected: %v", i+1, res, test.history)
		}
	}
}

func isEqual(a, b []models.CalcAction) bool {
	if a == nil && b == nil {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
