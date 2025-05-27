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
	return make([]models.CalcAction, 0)
}

func (m mockStorager) Set(uID string, action models.CalcAction) {}

var controller = newController(logrus.New(), mockCacher{}, mockStorager{})

func TestControllerSum(t *testing.T) {
	input := models.Input{
		Input: "1, 1",
	}

	res, err := controller.Sum("123", input)

	expected := models.CalcAction{
		Input:  input.Input,
		Action: models.Sum,
		Result: 2,
	}

	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if *res != expected {
		t.Errorf("Not equal!\nGot: %v\nExpected: %v\n", *res, expected)
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
	input := models.Input{
		Input: "1, 1",
	}
	res, err := controller.Mult("123", input)

	expected := models.CalcAction{
		Input:  input.Input,
		Action: models.Mult,
		Result: 1,
	}

	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if *res != expected {
		t.Errorf("Not equal!\nGot: %v\nExpected: %v\n", *res, expected)
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
